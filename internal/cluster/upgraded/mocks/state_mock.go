// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster/upgraded (interfaces: State)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	upgraded "github.com/spoke-d/thermionic/internal/cluster/upgraded"
	reflect "reflect"
)

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockState) Cluster() upgraded.Cluster {
	ret := m.ctrl.Call(m, "Cluster")
	ret0, _ := ret[0].(upgraded.Cluster)
	return ret0
}

// Cluster indicates an expected call of Cluster
func (mr *MockStateMockRecorder) Cluster() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockState)(nil).Cluster))
}

// Node mocks base method
func (m *MockState) Node() upgraded.Node {
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(upgraded.Node)
	return ret0
}

// Node indicates an expected call of Node
func (mr *MockStateMockRecorder) Node() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockState)(nil).Node))
}

// OS mocks base method
func (m *MockState) OS() upgraded.OS {
	ret := m.ctrl.Call(m, "OS")
	ret0, _ := ret[0].(upgraded.OS)
	return ret0
}

// OS indicates an expected call of OS
func (mr *MockStateMockRecorder) OS() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OS", reflect.TypeOf((*MockState)(nil).OS))
}
