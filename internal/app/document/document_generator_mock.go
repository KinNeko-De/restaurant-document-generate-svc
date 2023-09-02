// Code generated by mockery v2.33.0. DO NOT EDIT.

package document

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"

	v1 "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
)

// MockDocumentGenerator is an autogenerated mock type for the DocumentGenerator type
type MockDocumentGenerator struct {
	mock.Mock
}

type MockDocumentGenerator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDocumentGenerator) EXPECT() *MockDocumentGenerator_Expecter {
	return &MockDocumentGenerator_Expecter{mock: &_m.Mock}
}

// GenerateDocument provides a mock function with given fields: requestId, command
func (_m *MockDocumentGenerator) GenerateDocument(requestId uuid.UUID, command *v1.RequestedDocument) (GeneratedFile, error) {
	ret := _m.Called(requestId, command)

	var r0 GeneratedFile
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, *v1.RequestedDocument) (GeneratedFile, error)); ok {
		return rf(requestId, command)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, *v1.RequestedDocument) GeneratedFile); ok {
		r0 = rf(requestId, command)
	} else {
		r0 = ret.Get(0).(GeneratedFile)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, *v1.RequestedDocument) error); ok {
		r1 = rf(requestId, command)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDocumentGenerator_GenerateDocument_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateDocument'
type MockDocumentGenerator_GenerateDocument_Call struct {
	*mock.Call
}

// GenerateDocument is a helper method to define mock.On call
//   - requestId uuid.UUID
//   - command *v1.RequestedDocument
func (_e *MockDocumentGenerator_Expecter) GenerateDocument(requestId interface{}, command interface{}) *MockDocumentGenerator_GenerateDocument_Call {
	return &MockDocumentGenerator_GenerateDocument_Call{Call: _e.mock.On("GenerateDocument", requestId, command)}
}

func (_c *MockDocumentGenerator_GenerateDocument_Call) Run(run func(requestId uuid.UUID, command *v1.RequestedDocument)) *MockDocumentGenerator_GenerateDocument_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(*v1.RequestedDocument))
	})
	return _c
}

func (_c *MockDocumentGenerator_GenerateDocument_Call) Return(result GeneratedFile, err error) *MockDocumentGenerator_GenerateDocument_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *MockDocumentGenerator_GenerateDocument_Call) RunAndReturn(run func(uuid.UUID, *v1.RequestedDocument) (GeneratedFile, error)) *MockDocumentGenerator_GenerateDocument_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDocumentGenerator creates a new instance of MockDocumentGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDocumentGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDocumentGenerator {
	mock := &MockDocumentGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
