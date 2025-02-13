// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	cid "github.com/hyperledger/fabric-chaincode-go/v2/pkg/cid"

	mock "github.com/stretchr/testify/mock"

	shim "github.com/hyperledger/fabric-chaincode-go/v2/shim"
)

// TransactionContextInterface is an autogenerated mock type for the TransactionContextInterface type
type TransactionContextInterface struct {
	mock.Mock
}

// GetClientIdentity provides a mock function with given fields:
func (_m *TransactionContextInterface) GetClientIdentity() cid.ClientIdentity {
	ret := _m.Called()

	var r0 cid.ClientIdentity
	if rf, ok := ret.Get(0).(func() cid.ClientIdentity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cid.ClientIdentity)
		}
	}

	return r0
}

// GetStub provides a mock function with given fields:
func (_m *TransactionContextInterface) GetStub() shim.ChaincodeStubInterface {
	ret := _m.Called()

	var r0 shim.ChaincodeStubInterface
	if rf, ok := ret.Get(0).(func() shim.ChaincodeStubInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shim.ChaincodeStubInterface)
		}
	}

	return r0
}

type mockConstructorTestingTNewTransactionContextInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionContextInterface creates a new instance of TransactionContextInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionContextInterface(t mockConstructorTestingTNewTransactionContextInterface) *TransactionContextInterface {
	mock := &TransactionContextInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
