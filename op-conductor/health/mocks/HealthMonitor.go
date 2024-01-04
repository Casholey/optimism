// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// HealthMonitor is an autogenerated mock type for the HealthMonitor type
type HealthMonitor struct {
	mock.Mock
}

type HealthMonitor_Expecter struct {
	mock *mock.Mock
}

func (_m *HealthMonitor) EXPECT() *HealthMonitor_Expecter {
	return &HealthMonitor_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with given fields:
func (_m *HealthMonitor) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthMonitor_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type HealthMonitor_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *HealthMonitor_Expecter) Start() *HealthMonitor_Start_Call {
	return &HealthMonitor_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *HealthMonitor_Start_Call) Run(run func()) *HealthMonitor_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HealthMonitor_Start_Call) Return(_a0 error) *HealthMonitor_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HealthMonitor_Start_Call) RunAndReturn(run func() error) *HealthMonitor_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *HealthMonitor) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthMonitor_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type HealthMonitor_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *HealthMonitor_Expecter) Stop() *HealthMonitor_Stop_Call {
	return &HealthMonitor_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *HealthMonitor_Stop_Call) Run(run func()) *HealthMonitor_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HealthMonitor_Stop_Call) Return(_a0 error) *HealthMonitor_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HealthMonitor_Stop_Call) RunAndReturn(run func() error) *HealthMonitor_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Subscribe provides a mock function with given fields:
func (_m *HealthMonitor) Subscribe() <-chan bool {
	ret := _m.Called()

	var r0 <-chan bool
	if rf, ok := ret.Get(0).(func() <-chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan bool)
		}
	}

	return r0
}

// HealthMonitor_Subscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscribe'
type HealthMonitor_Subscribe_Call struct {
	*mock.Call
}

// Subscribe is a helper method to define mock.On call
func (_e *HealthMonitor_Expecter) Subscribe() *HealthMonitor_Subscribe_Call {
	return &HealthMonitor_Subscribe_Call{Call: _e.mock.On("Subscribe")}
}

func (_c *HealthMonitor_Subscribe_Call) Run(run func()) *HealthMonitor_Subscribe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HealthMonitor_Subscribe_Call) Return(_a0 <-chan bool) *HealthMonitor_Subscribe_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HealthMonitor_Subscribe_Call) RunAndReturn(run func() <-chan bool) *HealthMonitor_Subscribe_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewHealthMonitor interface {
	mock.TestingT
	Cleanup(func())
}

// NewHealthMonitor creates a new instance of HealthMonitor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHealthMonitor(t mockConstructorTestingTNewHealthMonitor) *HealthMonitor {
	mock := &HealthMonitor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
