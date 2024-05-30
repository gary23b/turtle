// Code generated by mockery v2.32.0. DO NOT EDIT.

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

// ClearScreen provides a mock function with given fields: c
func (_m *Canvas) ClearScreen(c color.RGBA) {
	_m.Called(c)
}

// CreateNewSprite provides a mock function with given fields:
func (_m *Canvas) CreateNewSprite() models.Sprite {
	ret := _m.Called()

	var r0 models.Sprite
	if rf, ok := ret.Get(0).(func() models.Sprite); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Sprite)
		}
	}

	return r0
}

// Exit provides a mock function with given fields:
func (_m *Canvas) Exit() {
	_m.Called()
}

// Fill provides a mock function with given fields: x, y, c
func (_m *Canvas) Fill(x int, y int, c color.RGBA) {
	_m.Called(x, y, c)
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

// PressedUserInput provides a mock function with given fields:
func (_m *Canvas) PressedUserInput() *models.UserInput {
	ret := _m.Called()

	var r0 *models.UserInput
	if rf, ok := ret.Get(0).(func() *models.UserInput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserInput)
		}
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

// SubscribeToJustPressedUserInput provides a mock function with given fields:
func (_m *Canvas) SubscribeToJustPressedUserInput() chan *models.UserInput {
	ret := _m.Called()

	var r0 chan *models.UserInput
	if rf, ok := ret.Get(0).(func() chan *models.UserInput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *models.UserInput)
		}
	}

	return r0
}

// UnSubscribeToJustPressedUserInput provides a mock function with given fields: in
func (_m *Canvas) UnSubscribeToJustPressedUserInput(in chan *models.UserInput) {
	_m.Called(in)
}

// NewCanvas creates a new instance of Canvas. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCanvas(t interface {
	mock.TestingT
	Cleanup(func())
}) *Canvas {
	mock := &Canvas{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
