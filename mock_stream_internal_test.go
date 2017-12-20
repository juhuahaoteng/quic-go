// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lucas-clemente/quic-go (interfaces: StreamI)

package quic

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	protocol "github.com/lucas-clemente/quic-go/internal/protocol"
	wire "github.com/lucas-clemente/quic-go/internal/wire"
	reflect "reflect"
	time "time"
)

// MockStreamI is a mock of StreamI interface
type MockStreamI struct {
	ctrl     *gomock.Controller
	recorder *MockStreamIMockRecorder
}

// MockStreamIMockRecorder is the mock recorder for MockStreamI
type MockStreamIMockRecorder struct {
	mock *MockStreamI
}

// NewMockStreamI creates a new mock instance
func NewMockStreamI(ctrl *gomock.Controller) *MockStreamI {
	mock := &MockStreamI{ctrl: ctrl}
	mock.recorder = &MockStreamIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStreamI) EXPECT() *MockStreamIMockRecorder {
	return _m.recorder
}

// CancelRead mocks base method
func (_m *MockStreamI) CancelRead(_param0 protocol.ApplicationErrorCode) error {
	ret := _m.ctrl.Call(_m, "CancelRead", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRead indicates an expected call of CancelRead
func (_mr *MockStreamIMockRecorder) CancelRead(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CancelRead", reflect.TypeOf((*MockStreamI)(nil).CancelRead), arg0)
}

// CancelWrite mocks base method
func (_m *MockStreamI) CancelWrite(_param0 protocol.ApplicationErrorCode) error {
	ret := _m.ctrl.Call(_m, "CancelWrite", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelWrite indicates an expected call of CancelWrite
func (_mr *MockStreamIMockRecorder) CancelWrite(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CancelWrite", reflect.TypeOf((*MockStreamI)(nil).CancelWrite), arg0)
}

// Close mocks base method
func (_m *MockStreamI) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockStreamIMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockStreamI)(nil).Close))
}

// Context mocks base method
func (_m *MockStreamI) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (_mr *MockStreamIMockRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Context", reflect.TypeOf((*MockStreamI)(nil).Context))
}

// Read mocks base method
func (_m *MockStreamI) Read(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Read", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (_mr *MockStreamIMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Read", reflect.TypeOf((*MockStreamI)(nil).Read), arg0)
}

// SetDeadline mocks base method
func (_m *MockStreamI) SetDeadline(_param0 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetDeadline", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline
func (_mr *MockStreamIMockRecorder) SetDeadline(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetDeadline", reflect.TypeOf((*MockStreamI)(nil).SetDeadline), arg0)
}

// SetReadDeadline mocks base method
func (_m *MockStreamI) SetReadDeadline(_param0 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetReadDeadline", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (_mr *MockStreamIMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetReadDeadline", reflect.TypeOf((*MockStreamI)(nil).SetReadDeadline), arg0)
}

// SetWriteDeadline mocks base method
func (_m *MockStreamI) SetWriteDeadline(_param0 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetWriteDeadline", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline
func (_mr *MockStreamIMockRecorder) SetWriteDeadline(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockStreamI)(nil).SetWriteDeadline), arg0)
}

// StreamID mocks base method
func (_m *MockStreamI) StreamID() protocol.StreamID {
	ret := _m.ctrl.Call(_m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID
func (_mr *MockStreamIMockRecorder) StreamID() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "StreamID", reflect.TypeOf((*MockStreamI)(nil).StreamID))
}

// Write mocks base method
func (_m *MockStreamI) Write(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Write", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (_mr *MockStreamIMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Write", reflect.TypeOf((*MockStreamI)(nil).Write), arg0)
}

// closeForShutdown mocks base method
func (_m *MockStreamI) closeForShutdown(_param0 error) {
	_m.ctrl.Call(_m, "closeForShutdown", _param0)
}

// closeForShutdown indicates an expected call of closeForShutdown
func (_mr *MockStreamIMockRecorder) closeForShutdown(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "closeForShutdown", reflect.TypeOf((*MockStreamI)(nil).closeForShutdown), arg0)
}

// finished mocks base method
func (_m *MockStreamI) finished() bool {
	ret := _m.ctrl.Call(_m, "finished")
	ret0, _ := ret[0].(bool)
	return ret0
}

// finished indicates an expected call of finished
func (_mr *MockStreamIMockRecorder) finished() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "finished", reflect.TypeOf((*MockStreamI)(nil).finished))
}

// getWindowUpdate mocks base method
func (_m *MockStreamI) getWindowUpdate() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "getWindowUpdate")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// getWindowUpdate indicates an expected call of getWindowUpdate
func (_mr *MockStreamIMockRecorder) getWindowUpdate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "getWindowUpdate", reflect.TypeOf((*MockStreamI)(nil).getWindowUpdate))
}

// handleMaxStreamDataFrame mocks base method
func (_m *MockStreamI) handleMaxStreamDataFrame(_param0 *wire.MaxStreamDataFrame) {
	_m.ctrl.Call(_m, "handleMaxStreamDataFrame", _param0)
}

// handleMaxStreamDataFrame indicates an expected call of handleMaxStreamDataFrame
func (_mr *MockStreamIMockRecorder) handleMaxStreamDataFrame(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "handleMaxStreamDataFrame", reflect.TypeOf((*MockStreamI)(nil).handleMaxStreamDataFrame), arg0)
}

// handleRstStreamFrame mocks base method
func (_m *MockStreamI) handleRstStreamFrame(_param0 *wire.RstStreamFrame) error {
	ret := _m.ctrl.Call(_m, "handleRstStreamFrame", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleRstStreamFrame indicates an expected call of handleRstStreamFrame
func (_mr *MockStreamIMockRecorder) handleRstStreamFrame(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "handleRstStreamFrame", reflect.TypeOf((*MockStreamI)(nil).handleRstStreamFrame), arg0)
}

// handleStopSendingFrame mocks base method
func (_m *MockStreamI) handleStopSendingFrame(_param0 *wire.StopSendingFrame) {
	_m.ctrl.Call(_m, "handleStopSendingFrame", _param0)
}

// handleStopSendingFrame indicates an expected call of handleStopSendingFrame
func (_mr *MockStreamIMockRecorder) handleStopSendingFrame(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "handleStopSendingFrame", reflect.TypeOf((*MockStreamI)(nil).handleStopSendingFrame), arg0)
}

// handleStreamFrame mocks base method
func (_m *MockStreamI) handleStreamFrame(_param0 *wire.StreamFrame) error {
	ret := _m.ctrl.Call(_m, "handleStreamFrame", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleStreamFrame indicates an expected call of handleStreamFrame
func (_mr *MockStreamIMockRecorder) handleStreamFrame(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "handleStreamFrame", reflect.TypeOf((*MockStreamI)(nil).handleStreamFrame), arg0)
}

// popStreamFrame mocks base method
func (_m *MockStreamI) popStreamFrame(_param0 protocol.ByteCount) *wire.StreamFrame {
	ret := _m.ctrl.Call(_m, "popStreamFrame", _param0)
	ret0, _ := ret[0].(*wire.StreamFrame)
	return ret0
}

// popStreamFrame indicates an expected call of popStreamFrame
func (_mr *MockStreamIMockRecorder) popStreamFrame(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "popStreamFrame", reflect.TypeOf((*MockStreamI)(nil).popStreamFrame), arg0)
}