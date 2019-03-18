// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster/events (interfaces: EventsSourceProvider,EventsSource)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	cert "github.com/spoke-d/thermionic/internal/cert"
	events "github.com/spoke-d/thermionic/internal/cluster/events"
	events0 "github.com/spoke-d/thermionic/pkg/events"
	reflect "reflect"
)

// MockEventsSourceProvider is a mock of EventsSourceProvider interface
type MockEventsSourceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockEventsSourceProviderMockRecorder
}

// MockEventsSourceProviderMockRecorder is the mock recorder for MockEventsSourceProvider
type MockEventsSourceProviderMockRecorder struct {
	mock *MockEventsSourceProvider
}

// NewMockEventsSourceProvider creates a new mock instance
func NewMockEventsSourceProvider(ctrl *gomock.Controller) *MockEventsSourceProvider {
	mock := &MockEventsSourceProvider{ctrl: ctrl}
	mock.recorder = &MockEventsSourceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventsSourceProvider) EXPECT() *MockEventsSourceProviderMockRecorder {
	return m.recorder
}

// Events mocks base method
func (m *MockEventsSourceProvider) Events(arg0 string, arg1 *cert.Info) (events.Client, events.EventsSource, error) {
	ret := m.ctrl.Call(m, "Events", arg0, arg1)
	ret0, _ := ret[0].(events.Client)
	ret1, _ := ret[1].(events.EventsSource)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Events indicates an expected call of Events
func (mr *MockEventsSourceProviderMockRecorder) Events(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockEventsSourceProvider)(nil).Events), arg0, arg1)
}

// MockEventsSource is a mock of EventsSource interface
type MockEventsSource struct {
	ctrl     *gomock.Controller
	recorder *MockEventsSourceMockRecorder
}

// MockEventsSourceMockRecorder is the mock recorder for MockEventsSource
type MockEventsSourceMockRecorder struct {
	mock *MockEventsSource
}

// NewMockEventsSource creates a new mock instance
func NewMockEventsSource(ctrl *gomock.Controller) *MockEventsSource {
	mock := &MockEventsSource{ctrl: ctrl}
	mock.recorder = &MockEventsSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventsSource) EXPECT() *MockEventsSourceMockRecorder {
	return m.recorder
}

// GetEvents mocks base method
func (m *MockEventsSource) GetEvents() (*events0.EventListener, error) {
	ret := m.ctrl.Call(m, "GetEvents")
	ret0, _ := ret[0].(*events0.EventListener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents
func (mr *MockEventsSourceMockRecorder) GetEvents() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockEventsSource)(nil).GetEvents))
}
