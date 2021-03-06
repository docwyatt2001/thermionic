// Code generated by MockGen. DO NOT EDIT.
// Source: net (interfaces: Listener)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	net "net"
	reflect "reflect"
)

// MockListener is a mock of Listener interface
type MockListener struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMockRecorder
}

// MockListenerMockRecorder is the mock recorder for MockListener
type MockListenerMockRecorder struct {
	mock *MockListener
}

// NewMockListener creates a new mock instance
func NewMockListener(ctrl *gomock.Controller) *MockListener {
	mock := &MockListener{ctrl: ctrl}
	mock.recorder = &MockListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListener) EXPECT() *MockListenerMockRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockListener) Accept() (net.Conn, error) {
	ret := m.ctrl.Call(m, "Accept")
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accept indicates an expected call of Accept
func (mr *MockListenerMockRecorder) Accept() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockListener)(nil).Accept))
}

// Addr mocks base method
func (m *MockListener) Addr() net.Addr {
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// Addr indicates an expected call of Addr
func (mr *MockListenerMockRecorder) Addr() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockListener)(nil).Addr))
}

// Close mocks base method
func (m *MockListener) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockListenerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockListener)(nil).Close))
}
