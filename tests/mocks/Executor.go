// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	api "github.com/microsoft/durabletask-go/api"
	backend "github.com/microsoft/durabletask-go/backend"

	context "context"

	mock "github.com/stretchr/testify/mock"

	protos "github.com/microsoft/durabletask-go/internal/protos"
)

// Executor is an autogenerated mock type for the Executor type
type Executor struct {
	mock.Mock
}

type Executor_Expecter struct {
	mock *mock.Mock
}

func (_m *Executor) EXPECT() *Executor_Expecter {
	return &Executor_Expecter{mock: &_m.Mock}
}

// ExecuteActivity provides a mock function with given fields: _a0, _a1, _a2
func (_m *Executor) ExecuteActivity(_a0 context.Context, _a1 api.InstanceID, _a2 *protos.HistoryEvent) (*protos.HistoryEvent, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteActivity")
	}

	var r0 *protos.HistoryEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, api.InstanceID, *protos.HistoryEvent) (*protos.HistoryEvent, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.InstanceID, *protos.HistoryEvent) *protos.HistoryEvent); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.HistoryEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.InstanceID, *protos.HistoryEvent) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Executor_ExecuteActivity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteActivity'
type Executor_ExecuteActivity_Call struct {
	*mock.Call
}

// ExecuteActivity is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 api.InstanceID
//   - _a2 *protos.HistoryEvent
func (_e *Executor_Expecter) ExecuteActivity(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Executor_ExecuteActivity_Call {
	return &Executor_ExecuteActivity_Call{Call: _e.mock.On("ExecuteActivity", _a0, _a1, _a2)}
}

func (_c *Executor_ExecuteActivity_Call) Run(run func(_a0 context.Context, _a1 api.InstanceID, _a2 *protos.HistoryEvent)) *Executor_ExecuteActivity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(api.InstanceID), args[2].(*protos.HistoryEvent))
	})
	return _c
}

func (_c *Executor_ExecuteActivity_Call) Return(_a0 *protos.HistoryEvent, _a1 error) *Executor_ExecuteActivity_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Executor_ExecuteActivity_Call) RunAndReturn(run func(context.Context, api.InstanceID, *protos.HistoryEvent) (*protos.HistoryEvent, error)) *Executor_ExecuteActivity_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteOrchestrator provides a mock function with given fields: ctx, iid, oldEvents, newEvents
func (_m *Executor) ExecuteOrchestrator(ctx context.Context, iid api.InstanceID, revision int, oldEvents []*protos.HistoryEvent, newEvents []*protos.HistoryEvent) (*backend.ExecutionResults, error) {
	ret := _m.Called(ctx, iid, revision, oldEvents, newEvents)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteOrchestrator")
	}

	var r0 *backend.ExecutionResults
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, api.InstanceID, int, []*protos.HistoryEvent, []*protos.HistoryEvent) (*backend.ExecutionResults, error)); ok {
		return rf(ctx, iid, revision, oldEvents, newEvents)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.InstanceID, int, []*protos.HistoryEvent, []*protos.HistoryEvent) *backend.ExecutionResults); ok {
		r0 = rf(ctx, iid, revision, oldEvents, newEvents)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backend.ExecutionResults)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.InstanceID, int, []*protos.HistoryEvent, []*protos.HistoryEvent) error); ok {
		r1 = rf(ctx, iid, revision, oldEvents, newEvents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Executor_ExecuteOrchestrator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteOrchestrator'
type Executor_ExecuteOrchestrator_Call struct {
	*mock.Call
}

// ExecuteOrchestrator is a helper method to define mock.On call
//   - ctx context.Context
//   - iid api.InstanceID
//   - oldEvents []*protos.HistoryEvent
//   - newEvents []*protos.HistoryEvent
func (_e *Executor_Expecter) ExecuteOrchestrator(ctx interface{}, iid interface{}, revision int, oldEvents interface{}, newEvents interface{}) *Executor_ExecuteOrchestrator_Call {
	return &Executor_ExecuteOrchestrator_Call{Call: _e.mock.On("ExecuteOrchestrator", ctx, iid, revision, oldEvents, newEvents)}
}

func (_c *Executor_ExecuteOrchestrator_Call) Run(run func(ctx context.Context, iid api.InstanceID, revision int, oldEvents []*protos.HistoryEvent, newEvents []*protos.HistoryEvent)) *Executor_ExecuteOrchestrator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(api.InstanceID), args[2].(int), args[3].([]*protos.HistoryEvent), args[4].([]*protos.HistoryEvent))
	})
	return _c
}

func (_c *Executor_ExecuteOrchestrator_Call) Return(_a0 *backend.ExecutionResults, _a1 error) *Executor_ExecuteOrchestrator_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Executor_ExecuteOrchestrator_Call) RunAndReturn(run func(context.Context, api.InstanceID, int, []*protos.HistoryEvent, []*protos.HistoryEvent) (*backend.ExecutionResults, error)) *Executor_ExecuteOrchestrator_Call {
	_c.Call.Return(run)
	return _c
}

// NewExecutor creates a new instance of Executor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Executor {
	mock := &Executor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
