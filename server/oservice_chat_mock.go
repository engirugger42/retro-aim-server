// Code generated by mockery v2.34.2. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mkaminski/goaim/oscar"
	mock "github.com/stretchr/testify/mock"

	state "github.com/mkaminski/goaim/state"
)

// mockOServiceChatHandler is an autogenerated mock type for the OServiceChatHandler type
type mockOServiceChatHandler struct {
	mock.Mock
}

type mockOServiceChatHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *mockOServiceChatHandler) EXPECT() *mockOServiceChatHandler_Expecter {
	return &mockOServiceChatHandler_Expecter{mock: &_m.Mock}
}

// ClientOnlineHandler provides a mock function with given fields: ctx, snacPayloadIn, sess, chatID
func (_m *mockOServiceChatHandler) ClientOnlineHandler(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x02_OServiceClientOnline, sess *state.Session, chatID string) error {
	ret := _m.Called(ctx, snacPayloadIn, sess, chatID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNAC_0x01_0x02_OServiceClientOnline, *state.Session, string) error); ok {
		r0 = rf(ctx, snacPayloadIn, sess, chatID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockOServiceChatHandler_ClientOnlineHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientOnlineHandler'
type mockOServiceChatHandler_ClientOnlineHandler_Call struct {
	*mock.Call
}

// ClientOnlineHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - snacPayloadIn oscar.SNAC_0x01_0x02_OServiceClientOnline
//   - sess *state.Session
//   - chatID string
func (_e *mockOServiceChatHandler_Expecter) ClientOnlineHandler(ctx interface{}, snacPayloadIn interface{}, sess interface{}, chatID interface{}) *mockOServiceChatHandler_ClientOnlineHandler_Call {
	return &mockOServiceChatHandler_ClientOnlineHandler_Call{Call: _e.mock.On("ClientOnlineHandler", ctx, snacPayloadIn, sess, chatID)}
}

func (_c *mockOServiceChatHandler_ClientOnlineHandler_Call) Run(run func(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x02_OServiceClientOnline, sess *state.Session, chatID string)) *mockOServiceChatHandler_ClientOnlineHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x01_0x02_OServiceClientOnline), args[2].(*state.Session), args[3].(string))
	})
	return _c
}

func (_c *mockOServiceChatHandler_ClientOnlineHandler_Call) Return(_a0 error) *mockOServiceChatHandler_ClientOnlineHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatHandler_ClientOnlineHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x01_0x02_OServiceClientOnline, *state.Session, string) error) *mockOServiceChatHandler_ClientOnlineHandler_Call {
	_c.Call.Return(run)
	return _c
}

// ClientVersionsHandler provides a mock function with given fields: ctx, snacPayloadIn
func (_m *mockOServiceChatHandler) ClientVersionsHandler(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions) oscar.XMessage {
	ret := _m.Called(ctx, snacPayloadIn)

	var r0 oscar.XMessage
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNAC_0x01_0x17_OServiceClientVersions) oscar.XMessage); ok {
		r0 = rf(ctx, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(oscar.XMessage)
	}

	return r0
}

// mockOServiceChatHandler_ClientVersionsHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientVersionsHandler'
type mockOServiceChatHandler_ClientVersionsHandler_Call struct {
	*mock.Call
}

// ClientVersionsHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions
func (_e *mockOServiceChatHandler_Expecter) ClientVersionsHandler(ctx interface{}, snacPayloadIn interface{}) *mockOServiceChatHandler_ClientVersionsHandler_Call {
	return &mockOServiceChatHandler_ClientVersionsHandler_Call{Call: _e.mock.On("ClientVersionsHandler", ctx, snacPayloadIn)}
}

func (_c *mockOServiceChatHandler_ClientVersionsHandler_Call) Run(run func(ctx context.Context, snacPayloadIn oscar.SNAC_0x01_0x17_OServiceClientVersions)) *mockOServiceChatHandler_ClientVersionsHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x01_0x17_OServiceClientVersions))
	})
	return _c
}

func (_c *mockOServiceChatHandler_ClientVersionsHandler_Call) Return(_a0 oscar.XMessage) *mockOServiceChatHandler_ClientVersionsHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatHandler_ClientVersionsHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x01_0x17_OServiceClientVersions) oscar.XMessage) *mockOServiceChatHandler_ClientVersionsHandler_Call {
	_c.Call.Return(run)
	return _c
}

// IdleNotificationHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *mockOServiceChatHandler) IdleNotificationHandler(ctx context.Context, sess *state.Session, snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification) error {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, oscar.SNAC_0x01_0x11_OServiceIdleNotification) error); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockOServiceChatHandler_IdleNotificationHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IdleNotificationHandler'
type mockOServiceChatHandler_IdleNotificationHandler_Call struct {
	*mock.Call
}

// IdleNotificationHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification
func (_e *mockOServiceChatHandler_Expecter) IdleNotificationHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *mockOServiceChatHandler_IdleNotificationHandler_Call {
	return &mockOServiceChatHandler_IdleNotificationHandler_Call{Call: _e.mock.On("IdleNotificationHandler", ctx, sess, snacPayloadIn)}
}

func (_c *mockOServiceChatHandler_IdleNotificationHandler_Call) Run(run func(ctx context.Context, sess *state.Session, snacPayloadIn oscar.SNAC_0x01_0x11_OServiceIdleNotification)) *mockOServiceChatHandler_IdleNotificationHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(oscar.SNAC_0x01_0x11_OServiceIdleNotification))
	})
	return _c
}

func (_c *mockOServiceChatHandler_IdleNotificationHandler_Call) Return(_a0 error) *mockOServiceChatHandler_IdleNotificationHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatHandler_IdleNotificationHandler_Call) RunAndReturn(run func(context.Context, *state.Session, oscar.SNAC_0x01_0x11_OServiceIdleNotification) error) *mockOServiceChatHandler_IdleNotificationHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsQueryHandler provides a mock function with given fields: ctx
func (_m *mockOServiceChatHandler) RateParamsQueryHandler(ctx context.Context) oscar.XMessage {
	ret := _m.Called(ctx)

	var r0 oscar.XMessage
	if rf, ok := ret.Get(0).(func(context.Context) oscar.XMessage); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(oscar.XMessage)
	}

	return r0
}

// mockOServiceChatHandler_RateParamsQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsQueryHandler'
type mockOServiceChatHandler_RateParamsQueryHandler_Call struct {
	*mock.Call
}

// RateParamsQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockOServiceChatHandler_Expecter) RateParamsQueryHandler(ctx interface{}) *mockOServiceChatHandler_RateParamsQueryHandler_Call {
	return &mockOServiceChatHandler_RateParamsQueryHandler_Call{Call: _e.mock.On("RateParamsQueryHandler", ctx)}
}

func (_c *mockOServiceChatHandler_RateParamsQueryHandler_Call) Run(run func(ctx context.Context)) *mockOServiceChatHandler_RateParamsQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockOServiceChatHandler_RateParamsQueryHandler_Call) Return(_a0 oscar.XMessage) *mockOServiceChatHandler_RateParamsQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatHandler_RateParamsQueryHandler_Call) RunAndReturn(run func(context.Context) oscar.XMessage) *mockOServiceChatHandler_RateParamsQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsSubAddHandler provides a mock function with given fields: _a0, _a1
func (_m *mockOServiceChatHandler) RateParamsSubAddHandler(_a0 context.Context, _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd) {
	_m.Called(_a0, _a1)
}

// mockOServiceChatHandler_RateParamsSubAddHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsSubAddHandler'
type mockOServiceChatHandler_RateParamsSubAddHandler_Call struct {
	*mock.Call
}

// RateParamsSubAddHandler is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd
func (_e *mockOServiceChatHandler_Expecter) RateParamsSubAddHandler(_a0 interface{}, _a1 interface{}) *mockOServiceChatHandler_RateParamsSubAddHandler_Call {
	return &mockOServiceChatHandler_RateParamsSubAddHandler_Call{Call: _e.mock.On("RateParamsSubAddHandler", _a0, _a1)}
}

func (_c *mockOServiceChatHandler_RateParamsSubAddHandler_Call) Run(run func(_a0 context.Context, _a1 oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *mockOServiceChatHandler_RateParamsSubAddHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd))
	})
	return _c
}

func (_c *mockOServiceChatHandler_RateParamsSubAddHandler_Call) Return() *mockOServiceChatHandler_RateParamsSubAddHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockOServiceChatHandler_RateParamsSubAddHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *mockOServiceChatHandler_RateParamsSubAddHandler_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserInfoFieldsHandler provides a mock function with given fields: ctx, sess, snacPayloadIn
func (_m *mockOServiceChatHandler) SetUserInfoFieldsHandler(ctx context.Context, sess *state.Session, snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (oscar.XMessage, error) {
	ret := _m.Called(ctx, sess, snacPayloadIn)

	var r0 oscar.XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (oscar.XMessage, error)); ok {
		return rf(ctx, sess, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) oscar.XMessage); ok {
		r0 = rf(ctx, sess, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(oscar.XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) error); ok {
		r1 = rf(ctx, sess, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockOServiceChatHandler_SetUserInfoFieldsHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserInfoFieldsHandler'
type mockOServiceChatHandler_SetUserInfoFieldsHandler_Call struct {
	*mock.Call
}

// SetUserInfoFieldsHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields
func (_e *mockOServiceChatHandler_Expecter) SetUserInfoFieldsHandler(ctx interface{}, sess interface{}, snacPayloadIn interface{}) *mockOServiceChatHandler_SetUserInfoFieldsHandler_Call {
	return &mockOServiceChatHandler_SetUserInfoFieldsHandler_Call{Call: _e.mock.On("SetUserInfoFieldsHandler", ctx, sess, snacPayloadIn)}
}

func (_c *mockOServiceChatHandler_SetUserInfoFieldsHandler_Call) Run(run func(ctx context.Context, sess *state.Session, snacPayloadIn oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields)) *mockOServiceChatHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields))
	})
	return _c
}

func (_c *mockOServiceChatHandler_SetUserInfoFieldsHandler_Call) Return(_a0 oscar.XMessage, _a1 error) *mockOServiceChatHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockOServiceChatHandler_SetUserInfoFieldsHandler_Call) RunAndReturn(run func(context.Context, *state.Session, oscar.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (oscar.XMessage, error)) *mockOServiceChatHandler_SetUserInfoFieldsHandler_Call {
	_c.Call.Return(run)
	return _c
}

// UserInfoQueryHandler provides a mock function with given fields: ctx, sess
func (_m *mockOServiceChatHandler) UserInfoQueryHandler(ctx context.Context, sess *state.Session) oscar.XMessage {
	ret := _m.Called(ctx, sess)

	var r0 oscar.XMessage
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session) oscar.XMessage); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Get(0).(oscar.XMessage)
	}

	return r0
}

// mockOServiceChatHandler_UserInfoQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserInfoQueryHandler'
type mockOServiceChatHandler_UserInfoQueryHandler_Call struct {
	*mock.Call
}

// UserInfoQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
func (_e *mockOServiceChatHandler_Expecter) UserInfoQueryHandler(ctx interface{}, sess interface{}) *mockOServiceChatHandler_UserInfoQueryHandler_Call {
	return &mockOServiceChatHandler_UserInfoQueryHandler_Call{Call: _e.mock.On("UserInfoQueryHandler", ctx, sess)}
}

func (_c *mockOServiceChatHandler_UserInfoQueryHandler_Call) Run(run func(ctx context.Context, sess *state.Session)) *mockOServiceChatHandler_UserInfoQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session))
	})
	return _c
}

func (_c *mockOServiceChatHandler_UserInfoQueryHandler_Call) Return(_a0 oscar.XMessage) *mockOServiceChatHandler_UserInfoQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatHandler_UserInfoQueryHandler_Call) RunAndReturn(run func(context.Context, *state.Session) oscar.XMessage) *mockOServiceChatHandler_UserInfoQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// WriteOServiceHostOnline provides a mock function with given fields:
func (_m *mockOServiceChatHandler) WriteOServiceHostOnline() oscar.XMessage {
	ret := _m.Called()

	var r0 oscar.XMessage
	if rf, ok := ret.Get(0).(func() oscar.XMessage); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(oscar.XMessage)
	}

	return r0
}

// mockOServiceChatHandler_WriteOServiceHostOnline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteOServiceHostOnline'
type mockOServiceChatHandler_WriteOServiceHostOnline_Call struct {
	*mock.Call
}

// WriteOServiceHostOnline is a helper method to define mock.On call
func (_e *mockOServiceChatHandler_Expecter) WriteOServiceHostOnline() *mockOServiceChatHandler_WriteOServiceHostOnline_Call {
	return &mockOServiceChatHandler_WriteOServiceHostOnline_Call{Call: _e.mock.On("WriteOServiceHostOnline")}
}

func (_c *mockOServiceChatHandler_WriteOServiceHostOnline_Call) Run(run func()) *mockOServiceChatHandler_WriteOServiceHostOnline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockOServiceChatHandler_WriteOServiceHostOnline_Call) Return(_a0 oscar.XMessage) *mockOServiceChatHandler_WriteOServiceHostOnline_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatHandler_WriteOServiceHostOnline_Call) RunAndReturn(run func() oscar.XMessage) *mockOServiceChatHandler_WriteOServiceHostOnline_Call {
	_c.Call.Return(run)
	return _c
}

// newMockOServiceChatHandler creates a new instance of mockOServiceChatHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockOServiceChatHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockOServiceChatHandler {
	mock := &mockOServiceChatHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
