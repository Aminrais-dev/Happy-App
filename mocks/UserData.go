// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	user "capstone/happyApp/features/user"

	mock "github.com/stretchr/testify/mock"
)

// DataUser is an autogenerated mock type for the DataInterface type
type DataUser struct {
	mock.Mock
}

// CheckStatus provides a mock function with given fields: _a0, _a1
func (_m *DataUser) CheckStatus(_a0 string, _a1 int) string {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, int) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CheckUsername provides a mock function with given fields: _a0
func (_m *DataUser) CheckUsername(_a0 string) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// DelUser provides a mock function with given fields: _a0
func (_m *DataUser) DelUser(_a0 int) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// InsertUser provides a mock function with given fields: _a0
func (_m *DataUser) InsertUser(_a0 user.CoreUser) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.CoreUser) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SelectUser provides a mock function with given fields: id
func (_m *DataUser) SelectUser(id int) (user.CoreUser, []user.CommunityProfile, error) {
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

// UpdtStatus provides a mock function with given fields: id, status
func (_m *DataUser) UpdtStatus(id int, status string) int {
	ret := _m.Called(id, status)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, string) int); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// UpdtUser provides a mock function with given fields: _a0
func (_m *DataUser) UpdtUser(_a0 user.CoreUser) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.CoreUser) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewDataUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataUser creates a new instance of DataUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataUser(t mockConstructorTestingTNewDataUser) *DataUser {
	mock := &DataUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
