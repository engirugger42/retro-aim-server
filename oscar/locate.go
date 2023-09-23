package oscar

import (
	"errors"
	"fmt"
	"io"
)

const (
	LocateErr                  uint16 = 0x0001
	LocateRightsQuery                 = 0x0002
	LocateRightsReply                 = 0x0003
	LocateSetInfo                     = 0x0004
	LocateUserInfoQuery               = 0x0005
	LocateUserInfoReply               = 0x0006
	LocateWatcherSubRequest           = 0x0007
	LocateWatcherNotification         = 0x0008
	LocateSetDirInfo                  = 0x0009
	LocateSetDirReply                 = 0x000A
	LocateGetDirInfo                  = 0x000B
	LocateGetDirReply                 = 0x000C
	LocateGroupCapabilityQuery        = 0x000D
	LocateGroupCapabilityReply        = 0x000E
	LocateSetKeywordInfo              = 0x000F
	LocateSetKeywordReply             = 0x0010
	LocateGetKeywordInfo              = 0x0011
	LocateGetKeywordReply             = 0x0012
	LocateFindListByEmail             = 0x0013
	LocateFindListReply               = 0x0014
	LocateUserInfoQuery2              = 0x0015
)

func routeLocate(sess *Session, sm *SessionManager, fm *FeedbagStore, snac snacFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	switch snac.subGroup {
	case LocateErr:
		panic("not implemented")
	case LocateRightsQuery:
		return SendAndReceiveLocateRights(snac, w, sequence)
	case LocateSetInfo:
		return ReceiveSetInfo(sess, sm, fm, snac, r)
	case LocateUserInfoQuery:
		panic("not implemented")
	case LocateUserInfoReply:
		panic("not implemented")
	case LocateWatcherSubRequest:
		panic("not implemented")
	case LocateWatcherNotification:
		panic("not implemented")
	case LocateSetDirInfo:
		return SendAndReceiveSetDirInfo(snac, r, w, sequence)
	case LocateGetDirInfo:
		return ReceiveLocateGetDirInfo(snac, r)
	case LocateGetDirReply:
		panic("not implemented")
	case LocateGroupCapabilityQuery:
		panic("not implemented")
	case LocateGroupCapabilityReply:
		panic("not implemented")
	case LocateSetKeywordInfo:
		return SendAndReceiveSetKeywordInfo(snac, r, w, sequence)
	case LocateGetKeywordInfo:
		panic("not implemented")
	case LocateGetKeywordReply:
		panic("not implemented")
	case LocateFindListByEmail:
		panic("not implemented")
	case LocateFindListReply:
		panic("not implemented")
	case LocateUserInfoQuery2:
		return SendAndReceiveUserInfoQuery2(sess, sm, fm, snac, r, w, sequence)
	}

	return nil
}

func SendAndReceiveLocateRights(snac snacFrame, w io.Writer, sequence *uint32) error {
	fmt.Printf("sendAndReceiveLocateRights read SNAC frame: %+v\n", snac)

	snacFrameOut := snacFrame{
		foodGroup: LOCATE,
		subGroup:  LocateRightsReply,
	}
	snacPayloadOut := SNAC_0x02_0x03_LocateRightsReply{
		TLVRestBlock: TLVRestBlock{
			TLVList: TLVList{
				{
					tType: 0x01,
					val:   uint16(1000),
				},
				{
					tType: 0x02,
					val:   uint16(1000),
				},
				{
					tType: 0x03,
					val:   uint16(1000),
				},
				{
					tType: 0x04,
					val:   uint16(1000),
				},
				{
					tType: 0x05,
					val:   uint16(1000),
				},
			},
		},
	}

	return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
}

var (
	LocateTlvTagsInfoSigMime         = uint16(0x01)
	LocateTlvTagsInfoSigData         = uint16(0x02)
	LocateTlvTagsInfoUnavailableMime = uint16(0x03)
	LocateTlvTagsInfoUnavailableData = uint16(0x04)
	LocateTlvTagsInfoCapabilities    = uint16(0x05)
	LocateTlvTagsInfoCerts           = uint16(0x06)
	LocateTlvTagsInfoSigTime         = uint16(0x0A)
	LocateTlvTagsInfoUnavailableTime = uint16(0x0B)
	LocateTlvTagsInfoSupportHostSig  = uint16(0x0C)
	LocateTlvTagsInfoHtmlInfoData    = uint16(0x0E)
	LocateTlvTagsInfoHtmlInfoType    = uint16(0x0D)
)

func ReceiveSetInfo(sess *Session, sm *SessionManager, fm *FeedbagStore, snac snacFrame, r io.Reader) error {
	fmt.Printf("ReceiveSetInfo read SNAC frame: %+v\n", snac)

	snacPayloadIn := SNAC_0x02_0x04_LocateSetInfo{}
	if err := Unmarshal(&snacPayloadIn, r); err != nil {
		return err
	}

	// update profile
	if profile, hasProfile := snacPayloadIn.getString(LocateTlvTagsInfoSigData); hasProfile {
		if err := fm.UpsertProfile(sess.ScreenName, profile); err != nil {
			return err
		}
	}

	// broadcast away message change to buddies
	if awayMsg, hasAwayMsg := snacPayloadIn.getString(LocateTlvTagsInfoUnavailableData); hasAwayMsg {
		sess.SetAwayMessage(awayMsg)
		if err := NotifyArrival(sess, sm, fm); err != nil {
			return err
		}
	}

	fmt.Printf("ReceiveSetInfo read SNAC: %+v\n", snacPayloadIn)

	return nil
}

func ReceiveLocateGetDirInfo(snac snacFrame, r io.Reader) error {
	fmt.Printf("ReceiveLocateGetDirInfo read SNAC frame: %+v\n", snac)

	snacPayloadIn := SNAC_0x02_0x0B_LocateGetDirInfo{}
	if err := Unmarshal(&snacPayloadIn, r); err != nil {
		return err
	}

	fmt.Printf("ReceiveLocateGetDirInfo read SNAC: %+v\n", snacPayloadIn)

	return nil
}

func SendAndReceiveUserInfoQuery2(sess *Session, sm *SessionManager, fm *FeedbagStore, snac snacFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	fmt.Printf("SendAndReceiveUserInfoQuery2 read SNAC frame: %+v\n", snac)

	snacPayloadIn := SNAC_0x02_0x15_LocateUserInfoQuery2{}
	if err := Unmarshal(&snacPayloadIn, r); err != nil {
		return err
	}

	blocked, err := fm.Blocked(sess.ScreenName, snacPayloadIn.ScreenName)
	if err != nil {
		return err
	}
	if blocked != BlockedNo {
		snacFrameOut := snacFrame{
			foodGroup: LOCATE,
			subGroup:  LocateErr,
		}
		snacPayloadOut := snacError{
			Code: ErrorCodeNotLoggedOn,
		}
		return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
	}

	buddySess, err := sm.RetrieveByScreenName(snacPayloadIn.ScreenName)
	if err != nil {
		if errors.Is(err, errSessNotFound) {
			snacFrameOut := snacFrame{
				foodGroup: LOCATE,
				subGroup:  LocateErr,
			}
			snacPayloadOut := snacError{
				Code: ErrorCodeNotLoggedOn,
			}
			return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
		}
	}

	snacFrameOut := snacFrame{
		foodGroup: LOCATE,
		subGroup:  LocateUserInfoReply,
	}
	snacPayloadOut := SNAC_0x02_0x06_LocateUserInfoReply{
		TLVUserInfo: TLVUserInfo{
			ScreenName:   snacPayloadIn.ScreenName,
			WarningLevel: buddySess.GetWarning(),
			TLVBlock: TLVBlock{
				TLVList: buddySess.GetUserInfo(),
			},
		},
		ClientProfile: TLVRestBlock{},
		AwayMessage:   TLVRestBlock{},
	}

	// profile
	if snacPayloadIn.Type2&1 == 1 {
		profile, err := fm.RetrieveProfile(snacPayloadIn.ScreenName)
		if err != nil {
			if err == errUserNotExist {
				snacFrameOut := snacFrame{
					foodGroup: LOCATE,
					subGroup:  LocateErr,
				}
				snacPayloadOut := snacError{
					Code: ErrorCodeNotLoggedOn,
				}
				return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
			}
			return err
		}
		snacPayloadOut.ClientProfile.TLVList = append(snacPayloadOut.ClientProfile.TLVList, []TLV{
			{
				tType: 0x01,
				val:   `text/aolrtf; charset="us-ascii"`,
			},
			{
				tType: 0x02,
				val:   profile,
			},
		}...)
	}

	// away message
	if snacPayloadIn.Type2&2 == 2 {
		snacPayloadOut.ClientProfile.TLVList = append(snacPayloadOut.ClientProfile.TLVList, []TLV{
			{
				tType: 0x03,
				val:   `text/aolrtf; charset="us-ascii"`,
			},
			{
				tType: 0x04,
				val:   buddySess.GetAwayMessage(),
			},
		}...)
	}

	return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
}

func SendAndReceiveSetDirInfo(snac snacFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	fmt.Printf("SendAndReceiveSetDirInfo read SNAC frame: %+v\n", snac)

	snacPayloadIn := SNAC_0x02_0x09_LocateSetDirInfo{}
	if err := Unmarshal(&snacPayloadIn, r); err != nil {
		return err
	}

	snacFrameOut := snacFrame{
		foodGroup: LOCATE,
		subGroup:  LocateSetDirReply,
	}
	snacPayloadOut := SNAC_0x02_0x0A_LocateSetDirReply{
		Result: 1,
	}

	return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
}

func SendAndReceiveSetKeywordInfo(snac snacFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	fmt.Printf("SendAndReceiveSetKeywordInfo read SNAC frame: %+v\n", snac)

	snacPayloadIn := SNAC_0x02_0x0F_LocateSetKeywordInfo{}
	if err := Unmarshal(&snacPayloadIn, r); err != nil {
		return err
	}

	snacFrameOut := snacFrame{
		foodGroup: LOCATE,
		subGroup:  LocateSetKeywordReply,
	}
	snacPayloadOut := SNAC_0x02_0x10_LocateSetKeywordReply{
		Unknown: 1,
	}

	return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
}
