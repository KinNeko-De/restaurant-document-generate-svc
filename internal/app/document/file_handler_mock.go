// Code generated by mockery v2.33.0. DO NOT EDIT.

package document

import mock "github.com/stretchr/testify/mock"

// MockFileHandler is an autogenerated mock type for the FileHandler type
type MockFileHandler struct {
	mock.Mock
}

type MockFileHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFileHandler) EXPECT() *MockFileHandler_Expecter {
	return &MockFileHandler_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockFileHandler) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileHandler_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockFileHandler_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockFileHandler_Expecter) Close() *MockFileHandler_Close_Call {
	return &MockFileHandler_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockFileHandler_Close_Call) Run(run func()) *MockFileHandler_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileHandler_Close_Call) Return(_a0 error) *MockFileHandler_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileHandler_Close_Call) RunAndReturn(run func() error) *MockFileHandler_Close_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFileHandler creates a new instance of MockFileHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFileHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFileHandler {
	mock := &MockFileHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
