// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	queryresult "github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	mock "github.com/stretchr/testify/mock"
)

// HistoryQueryIteratorInterface is an autogenerated mock type for the HistoryQueryIteratorInterface type
type HistoryQueryIteratorInterface struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *HistoryQueryIteratorInterface) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasNext provides a mock function with given fields:
func (_m *HistoryQueryIteratorInterface) HasNext() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *HistoryQueryIteratorInterface) Next() (*queryresult.KeyModification, error) {
	ret := _m.Called()

	var r0 *queryresult.KeyModification
	if rf, ok := ret.Get(0).(func() *queryresult.KeyModification); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*queryresult.KeyModification)
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

type mockConstructorTestingTNewHistoryQueryIteratorInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewHistoryQueryIteratorInterface creates a new instance of HistoryQueryIteratorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHistoryQueryIteratorInterface(t mockConstructorTestingTNewHistoryQueryIteratorInterface) *HistoryQueryIteratorInterface {
	mock := &HistoryQueryIteratorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
