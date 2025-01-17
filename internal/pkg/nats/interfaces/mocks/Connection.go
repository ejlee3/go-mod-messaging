// Code generated by mockery v2.13.1. DO NOT EDIT.
// note that nats.go will need to be added to the end of the nats import
package mocks

import (
	nats "github.com/nats-io/nats.go"
	mock "github.com/stretchr/testify/mock"
)

// Connection is an autogenerated mock type for the Connection type
type Connection struct {
	mock.Mock
}

// Drain provides a mock function with given fields:
func (_m *Connection) Drain() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishMsg provides a mock function with given fields: _a0
func (_m *Connection) PublishMsg(_a0 *nats.Msg) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*nats.Msg) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueSubscribe provides a mock function with given fields: _a0, _a1, _a2
func (_m *Connection) QueueSubscribe(_a0 string, _a1 string, _a2 nats.MsgHandler) (*nats.Subscription, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *nats.Subscription
	if rf, ok := ret.Get(0).(func(string, string, nats.MsgHandler) *nats.Subscription); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nats.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, nats.MsgHandler) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewConnection interface {
	mock.TestingT
	Cleanup(func())
}

// NewConnection creates a new instance of Connection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConnection(t mockConstructorTestingTNewConnection) *Connection {
	mock := &Connection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
