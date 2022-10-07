// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	event "capstone/happyApp/features/event"

	coreapi "github.com/midtrans/midtrans-go/coreapi"

	mock "github.com/stretchr/testify/mock"
)

// DataEvent is an autogenerated mock type for the DataInterface type
type DataEvent struct {
	mock.Mock
}

// CreatePayment provides a mock function with given fields: reqMidtrans, userId, EventId, method
func (_m *DataEvent) CreatePayment(reqMidtrans coreapi.ChargeReq, userId int, EventId int, method string) (*coreapi.ChargeResponse, error) {
	ret := _m.Called(reqMidtrans, userId, EventId, method)

	var r0 *coreapi.ChargeResponse
	if rf, ok := ret.Get(0).(func(coreapi.ChargeReq, int, int, string) *coreapi.ChargeResponse); ok {
		r0 = rf(reqMidtrans, userId, EventId, method)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coreapi.ChargeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(coreapi.ChargeReq, int, int, string) error); ok {
		r1 = rf(reqMidtrans, userId, EventId, method)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembers provides a mock function with given fields: _a0
func (_m *DataEvent) GetMembers(_a0 []event.Response) []event.Response {
	ret := _m.Called(_a0)

	var r0 []event.Response
	if rf, ok := ret.Get(0).(func([]event.Response) []event.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]event.Response)
		}
	}

	return r0
}

// InsertEvent provides a mock function with given fields: _a0, _a1
func (_m *DataEvent) InsertEvent(_a0 event.EventCore, _a1 int) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(event.EventCore, int) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SelectAmountEvent provides a mock function with given fields: idEvent
func (_m *DataEvent) SelectAmountEvent(idEvent int) uint64 {
	ret := _m.Called(idEvent)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(int) uint64); ok {
		r0 = rf(idEvent)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// SelectEvent provides a mock function with given fields: _a0
func (_m *DataEvent) SelectEvent(_a0 string) ([]event.Response, error) {
	ret := _m.Called(_a0)

	var r0 []event.Response
	if rf, ok := ret.Get(0).(func(string) []event.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]event.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectEventComu provides a mock function with given fields: data, idComu, userId
func (_m *DataEvent) SelectEventComu(data []event.Response, idComu int, userId int) (event.CommunityEvent, error) {
	ret := _m.Called(data, idComu, userId)

	var r0 event.CommunityEvent
	if rf, ok := ret.Get(0).(func([]event.Response, int, int) event.CommunityEvent); ok {
		r0 = rf(data, idComu, userId)
	} else {
		r0 = ret.Get(0).(event.CommunityEvent)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]event.Response, int, int) error); ok {
		r1 = rf(data, idComu, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectEventDetail provides a mock function with given fields: idEvent, userId
func (_m *DataEvent) SelectEventDetail(idEvent int, userId int) (event.EventDetail, error) {
	ret := _m.Called(idEvent, userId)

	var r0 event.EventDetail
	if rf, ok := ret.Get(0).(func(int, int) event.EventDetail); ok {
		r0 = rf(idEvent, userId)
	} else {
		r0 = ret.Get(0).(event.EventDetail)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(idEvent, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDataEvent interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataEvent creates a new instance of DataEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataEvent(t mockConstructorTestingTNewDataEvent) *DataEvent {
	mock := &DataEvent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
