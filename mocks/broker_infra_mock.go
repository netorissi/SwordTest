// Code generated by MockGen. DO NOT EDIT.
// Source: ./broker.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nats "github.com/nats-io/nats.go"
)

// MockMessageBroker is a mock of MessageBroker interface.
type MockMessageBroker struct {
	ctrl     *gomock.Controller
	recorder *MockMessageBrokerMockRecorder
}

// MockMessageBrokerMockRecorder is the mock recorder for MockMessageBroker.
type MockMessageBrokerMockRecorder struct {
	mock *MockMessageBroker
}

// NewMockMessageBroker creates a new mock instance.
func NewMockMessageBroker(ctrl *gomock.Controller) *MockMessageBroker {
	mock := &MockMessageBroker{ctrl: ctrl}
	mock.recorder = &MockMessageBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageBroker) EXPECT() *MockMessageBrokerMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockMessageBroker) Publish(topic string, msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", topic, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockMessageBrokerMockRecorder) Publish(topic, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockMessageBroker)(nil).Publish), topic, msg)
}

// Subscribe mocks base method.
func (m *MockMessageBroker) Subscribe(topic string, f func(*nats.Msg)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", topic, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockMessageBrokerMockRecorder) Subscribe(topic, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockMessageBroker)(nil).Subscribe), topic, f)
}
