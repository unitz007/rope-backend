// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Server is an autogenerated mock type for the Server type
type Server struct {
	mock.Mock
}

// AttachMiddleware provides a mock function with given fields: middleware
func (_m *Server) AttachMiddleware(middleware func(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request)) Server {
	ret := _m.Called(middleware)

	if len(ret) == 0 {
		panic("no return value specified for AttachMiddleware")
	}

	var r0 Server
	if rf, ok := ret.Get(0).(func(func(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request)) Server); ok {
		r0 = rf(middleware)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Server)
		}
	}

	return r0
}

// RegisterHandler provides a mock function with given fields: path, method, handler
func (_m *Server) RegisterHandler(path string, method string, handler func(http.ResponseWriter, *http.Request)) {
	_m.Called(path, method, handler)
}

// Start provides a mock function with given fields: address
func (_m *Server) Start(address string) error {
	ret := _m.Called(address)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServer creates a new instance of Server. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Server {
	mock := &Server{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
