// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	color "image/color"

	models "github.com/gary23b/turtle/models"
	mock "github.com/stretchr/testify/mock"
)

// Canvas is an autogenerated mock type for the Canvas type
type Canvas struct {
	mock.Mock
}

// FillScreen provides a mock function with given fields: c
func (_m *Canvas) FillScreen(c color.RGBA) {
	_m.Called(c)
}

// GetHeight provides a mock function with given fields:
func (_m *Canvas) GetHeight() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetUserInput provides a mock function with given fields:
func (_m *Canvas) GetUserInput() models.UserInput {
	ret := _m.Called()

	var r0 models.UserInput
	if rf, ok := ret.Get(0).(func() models.UserInput); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.UserInput)
	}

	return r0
}

// GetWidth provides a mock function with given fields:
func (_m *Canvas) GetWidth() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SetCartesianPixel provides a mock function with given fields: x, y, c
func (_m *Canvas) SetCartesianPixel(x int, y int, c color.RGBA) {
	_m.Called(x, y, c)
}

// SetPixel provides a mock function with given fields: x, y, c
func (_m *Canvas) SetPixel(x int, y int, c color.RGBA) {
	_m.Called(x, y, c)
}

type mockConstructorTestingTNewCanvas interface {
	mock.TestingT
	Cleanup(func())
}

// NewCanvas creates a new instance of Canvas. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCanvas(t mockConstructorTestingTNewCanvas) *Canvas {
	mock := &Canvas{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
