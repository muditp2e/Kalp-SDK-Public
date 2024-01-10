// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	x509 "crypto/x509"

	mock "github.com/stretchr/testify/mock"
)

// ClientIdentity is an autogenerated mock type for the ClientIdentity type
type ClientIdentity struct {
	mock.Mock
}

// AssertAttributeValue provides a mock function with given fields: attrName, attrValue
func (_m *ClientIdentity) AssertAttributeValue(attrName string, attrValue string) error {
	ret := _m.Called(attrName, attrValue)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(attrName, attrValue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAttributeValue provides a mock function with given fields: attrName
func (_m *ClientIdentity) GetAttributeValue(attrName string) (string, bool, error) {
	ret := _m.Called(attrName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(attrName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(attrName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(attrName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetID provides a mock function with given fields:
func (_m *ClientIdentity) GetID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMSPID provides a mock function with given fields:
func (_m *ClientIdentity) GetMSPID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetX509Certificate provides a mock function with given fields:
func (_m *ClientIdentity) GetX509Certificate() (*x509.Certificate, error) {
	ret := _m.Called()

	var r0 *x509.Certificate
	if rf, ok := ret.Get(0).(func() *x509.Certificate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*x509.Certificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClientIdentity interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientIdentity creates a new instance of ClientIdentity. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientIdentity(t mockConstructorTestingTNewClientIdentity) *ClientIdentity {
	mock := &ClientIdentity{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}