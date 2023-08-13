// Code generated by mockery v2.32.0. DO NOT EDIT.

package document

import (
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"

	v1 "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
)

// DocumentGenerator is an autogenerated mock type for the DocumentGenerator type
type DocumentGeneratorMock struct {
	mock.Mock
}

type DocumentGenerator_Expecter struct {
	mock *mock.Mock
}

func (_m *DocumentGeneratorMock) EXPECT() *DocumentGenerator_Expecter {
	return &DocumentGenerator_Expecter{mock: &_m.Mock}
}

// GenerateDocument provides a mock function with given fields: requestId, command
func (_m *DocumentGeneratorMock) GenerateDocument(requestId uuid.UUID, command *v1.RequestedDocument) (GenerationResult, error) {
	ret := _m.Called(requestId, command)

	var r0 GenerationResult
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, *v1.RequestedDocument) (GenerationResult, error)); ok {
		return rf(requestId, command)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, *v1.RequestedDocument) GenerationResult); ok {
		r0 = rf(requestId, command)
	} else {
		r0 = ret.Get(0).(GenerationResult)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, *v1.RequestedDocument) error); ok {
		r1 = rf(requestId, command)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DocumentGenerator_GenerateDocument_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateDocument'
type DocumentGenerator_GenerateDocument_Call struct {
	*mock.Call
}

// GenerateDocument is a helper method to define mock.On call
//   - requestId uuid.UUID
//   - command *v1.RequestedDocument
func (_e *DocumentGenerator_Expecter) GenerateDocument(requestId interface{}, command interface{}) *DocumentGenerator_GenerateDocument_Call {
	return &DocumentGenerator_GenerateDocument_Call{Call: _e.mock.On("GenerateDocument", requestId, command)}
}

func (_c *DocumentGenerator_GenerateDocument_Call) Run(run func(requestId uuid.UUID, command *v1.RequestedDocument)) *DocumentGenerator_GenerateDocument_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(*v1.RequestedDocument))
	})
	return _c
}

func (_c *DocumentGenerator_GenerateDocument_Call) Return(result GenerationResult, err error) *DocumentGenerator_GenerateDocument_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *DocumentGenerator_GenerateDocument_Call) RunAndReturn(run func(uuid.UUID, *v1.RequestedDocument) (GenerationResult, error)) *DocumentGenerator_GenerateDocument_Call {
	_c.Call.Return(run)
	return _c
}

// NewDocumentGenerator creates a new instance of DocumentGeneratorMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDocumentGeneratorMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DocumentGeneratorMock {
	mock := &DocumentGeneratorMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}