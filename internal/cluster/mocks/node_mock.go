// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster (interfaces: Node)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	db "github.com/spoke-d/thermionic/internal/db"
	database "github.com/spoke-d/thermionic/internal/db/database"
	reflect "reflect"
)

// MockNode is a mock of Node interface
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockNode) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockNodeMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNode)(nil).Close))
}

// DB mocks base method
func (m *MockNode) DB() database.DB {
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(database.DB)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockNodeMockRecorder) DB() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockNode)(nil).DB))
}

// Dir mocks base method
func (m *MockNode) Dir() string {
	ret := m.ctrl.Call(m, "Dir")
	ret0, _ := ret[0].(string)
	return ret0
}

// Dir indicates an expected call of Dir
func (mr *MockNodeMockRecorder) Dir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dir", reflect.TypeOf((*MockNode)(nil).Dir))
}

// Open mocks base method
func (m *MockNode) Open(arg0 string, arg1 func(*db.Node) error) error {
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockNodeMockRecorder) Open(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockNode)(nil).Open), arg0, arg1)
}

// Transaction mocks base method
func (m *MockNode) Transaction(arg0 func(*db.NodeTx) error) error {
	ret := m.ctrl.Call(m, "Transaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction
func (mr *MockNodeMockRecorder) Transaction(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockNode)(nil).Transaction), arg0)
}
