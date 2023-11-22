package server

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"

	"github.com/mkaminski/goaim/oscar"
	"github.com/mkaminski/goaim/state"
)

var (
	ErrUnsupportedSubGroup = errors.New("unimplemented subgroup, your client version may be unsupported")
)

type (
	incomingMessage struct {
		flap    oscar.FlapFrame
		payload *bytes.Buffer
	}
	alertHandler     func(ctx context.Context, msg oscar.XMessage, w io.Writer, u *uint32) error
	clientReqHandler func(ctx context.Context, r io.Reader, w io.Writer, u *uint32) error
)

func sendSNAC(originsnac oscar.SnacFrame, snacFrame oscar.SnacFrame, snacOut any, sequence *uint32, w io.Writer) error {
	if originsnac.RequestID != 0 {
		snacFrame.RequestID = originsnac.RequestID
	}

	snacBuf := &bytes.Buffer{}
	if err := oscar.Marshal(snacFrame, snacBuf); err != nil {
		return err
	}
	if err := oscar.Marshal(snacOut, snacBuf); err != nil {
		return err
	}

	flap := oscar.FlapFrame{
		StartMarker:   42,
		FrameType:     oscar.FlapFrameData,
		Sequence:      uint16(*sequence),
		PayloadLength: uint16(snacBuf.Len()),
	}

	if err := oscar.Marshal(flap, w); err != nil {
		return err
	}

	expectLen := snacBuf.Len()
	c, err := w.Write(snacBuf.Bytes())
	if err != nil {
		return err
	}
	if c != expectLen {
		panic("did not write the expected # of bytes")
	}

	*sequence++
	return nil
}

func sendInvalidSNACErr(snac oscar.SnacFrame, w io.Writer, sequence *uint32) error {
	snacFrameOut := oscar.SnacFrame{
		FoodGroup: snac.FoodGroup,
		SubGroup:  0x01, // error subgroup for all SNACs
	}
	snacPayloadOut := oscar.SnacError{
		Code: oscar.ErrorCodeInvalidSnac,
	}
	return sendSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
}

func consumeFLAPFrames(r io.Reader, msgCh chan incomingMessage, errCh chan error) {
	defer close(msgCh)
	defer close(errCh)

	for {
		in := incomingMessage{}
		if err := oscar.Unmarshal(&in.flap, r); err != nil {
			errCh <- err
			return
		}

		if in.flap.FrameType == oscar.FlapFrameData {
			buf := make([]byte, in.flap.PayloadLength)
			if _, err := r.Read(buf); err != nil {
				errCh <- err
				return
			}
			in.payload = bytes.NewBuffer(buf)
		}

		msgCh <- in
	}
}

func dispatchIncomingMessages(ctx context.Context, sess *state.Session, seq uint32, rw io.ReadWriter, logger *slog.Logger, fn clientReqHandler, alertHandler alertHandler) {
	// buffered so that the go routine has room to exit
	msgCh := make(chan incomingMessage, 1)
	readErrCh := make(chan error, 1)
	go consumeFLAPFrames(rw, msgCh, readErrCh)

	defer func() {
		logger.InfoContext(ctx, "user disconnected")
	}()

	for {
		select {
		case m := <-msgCh:
			switch m.flap.FrameType {
			case oscar.FlapFrameData:
				// route a client request to the appropriate service handler. the
				// handler may write a response to the client connection.
				if err := fn(ctx, m.payload, rw, &seq); err != nil {
					return
				}
			case oscar.FlapFrameSignon:
				logger.ErrorContext(ctx, "shouldn't get FlapFrameSignon", "flap", m.flap)
			case oscar.FlapFrameError:
				logger.ErrorContext(ctx, "got FlapFrameError", "flap", m.flap)
				return
			case oscar.FlapFrameSignoff:
				logger.InfoContext(ctx, "got FlapFrameSignoff", "flap", m.flap)
				return
			case oscar.FlapFrameKeepAlive:
				logger.DebugContext(ctx, "keepalive heartbeat")
			default:
				logger.ErrorContext(ctx, "got unknown FLAP frame type", "flap", m.flap)
				return
			}
		case m := <-sess.RecvMessage():
			// forward a notification sent from another client to this client
			if err := alertHandler(ctx, m, rw, &seq); err != nil {
				logRequestError(ctx, logger, m.SnacFrame, err)
				return
			}
			logRequest(ctx, logger, m.SnacFrame, m.SnacOut)
		case <-sess.Closed():
			// gracefully disconnect so that the client does not try to
			// reconnect when the connection closes.
			flap := oscar.FlapFrame{
				StartMarker:   42,
				FrameType:     oscar.FlapFrameSignoff,
				Sequence:      uint16(seq),
				PayloadLength: uint16(0),
			}
			if err := oscar.Marshal(flap, rw); err != nil {
				logger.ErrorContext(ctx, "unable to gracefully disconnect user", "err", err)
			}
			return
		case err := <-readErrCh:
			if !errors.Is(io.EOF, err) {
				logger.ErrorContext(ctx, "client disconnected with error", "err", err)
			}
			return
		}
	}
}
