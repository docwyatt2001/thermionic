// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster/raft (interfaces: RaftProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	raft0 "github.com/hashicorp/raft"
	raft "github.com/spoke-d/thermionic/internal/cluster/raft"
	reflect "reflect"
)

// MockRaftProvider is a mock of RaftProvider interface
type MockRaftProvider struct {
	ctrl     *gomock.Controller
	recorder *MockRaftProviderMockRecorder
}

// MockRaftProviderMockRecorder is the mock recorder for MockRaftProvider
type MockRaftProviderMockRecorder struct {
	mock *MockRaftProvider
}

// NewMockRaftProvider creates a new mock instance
func NewMockRaftProvider(ctrl *gomock.Controller) *MockRaftProvider {
	mock := &MockRaftProvider{ctrl: ctrl}
	mock.recorder = &MockRaftProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRaftProvider) EXPECT() *MockRaftProviderMockRecorder {
	return m.recorder
}

// Raft mocks base method
func (m *MockRaftProvider) Raft(arg0 *raft0.Config, arg1 raft0.FSM, arg2 raft.LogStore, arg3 raft0.SnapshotStore, arg4 raft0.Transport) (*raft0.Raft, error) {
	ret := m.ctrl.Call(m, "Raft", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*raft0.Raft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Raft indicates an expected call of Raft
func (mr *MockRaftProviderMockRecorder) Raft(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raft", reflect.TypeOf((*MockRaftProvider)(nil).Raft), arg0, arg1, arg2, arg3, arg4)
}
