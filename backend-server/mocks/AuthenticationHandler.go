// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mux "github.com/gorilla/mux"
	mock "github.com/stretchr/testify/mock"
)

// AuthenticationHandler is an autogenerated mock type for the AuthenticationHandler type
type AuthenticationHandler struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: w, r
func (_m *AuthenticationHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetMux provides a mock function with given fields:
func (_m *AuthenticationHandler) GetMux() *mux.Router {
	ret := _m.Called()

	var r0 *mux.Router
	if rf, ok := ret.Get(0).(func() *mux.Router); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mux.Router)
		}
	}

	return r0
}

// LoginPage provides a mock function with given fields: w, r
func (_m *AuthenticationHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

type mockConstructorTestingTNewAuthenticationHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthenticationHandler creates a new instance of AuthenticationHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthenticationHandler(t mockConstructorTestingTNewAuthenticationHandler) *AuthenticationHandler {
	mock := &AuthenticationHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
