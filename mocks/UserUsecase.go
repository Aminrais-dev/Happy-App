// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	user "capstone/happyApp/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UsecaseUser is an autogenerated mock type for the UsecaseInterface type
type UsecaseUser struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: _a0
func (_m *UsecaseUser) DeleteUser(_a0 int) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetUser provides a mock function with given fields: id
func (_m *UsecaseUser) GetUser(id int) (user.CoreUser, []user.CommunityProfile, error) {
	ret := _m.Called(id)

	var r0 user.CoreUser
	if rf, ok := ret.Get(0).(func(int) user.CoreUser); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.CoreUser)
	}

	var r1 []user.CommunityProfile
	if rf, ok := ret.Get(1).(func(int) []user.CommunityProfile); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]user.CommunityProfile)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PostUser provides a mock function with given fields: _a0
func (_m *UsecaseUser) PostUser(_a0 user.CoreUser) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.CoreUser) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: id
func (_m *UsecaseUser) UpdateStatus(id int) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *UsecaseUser) UpdateUser(_a0 user.CoreUser) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.CoreUser) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewUsecaseUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecaseUser creates a new instance of UsecaseUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecaseUser(t mockConstructorTestingTNewUsecaseUser) *UsecaseUser {
	mock := &UsecaseUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}