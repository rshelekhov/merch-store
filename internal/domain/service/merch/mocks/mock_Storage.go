// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/rshelekhov/avito-tech-internship/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

type Storage_Expecter struct {
	mock *mock.Mock
}

func (_m *Storage) EXPECT() *Storage_Expecter {
	return &Storage_Expecter{mock: &_m.Mock}
}

// AddToInventory provides a mock function with given fields: ctx, purchase
func (_m *Storage) AddToInventory(ctx context.Context, purchase entity.Purchase) error {
	ret := _m.Called(ctx, purchase)

	if len(ret) == 0 {
		panic("no return value specified for AddToInventory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Purchase) error); ok {
		r0 = rf(ctx, purchase)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_AddToInventory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddToInventory'
type Storage_AddToInventory_Call struct {
	*mock.Call
}

// AddToInventory is a helper method to define mock.On call
//   - ctx context.Context
//   - purchase entity.Purchase
func (_e *Storage_Expecter) AddToInventory(ctx interface{}, purchase interface{}) *Storage_AddToInventory_Call {
	return &Storage_AddToInventory_Call{Call: _e.mock.On("AddToInventory", ctx, purchase)}
}

func (_c *Storage_AddToInventory_Call) Run(run func(ctx context.Context, purchase entity.Purchase)) *Storage_AddToInventory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Purchase))
	})
	return _c
}

func (_c *Storage_AddToInventory_Call) Return(_a0 error) *Storage_AddToInventory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_AddToInventory_Call) RunAndReturn(run func(context.Context, entity.Purchase) error) *Storage_AddToInventory_Call {
	_c.Call.Return(run)
	return _c
}

// GetMerchByName provides a mock function with given fields: ctx, itemName
func (_m *Storage) GetMerchByName(ctx context.Context, itemName string) (entity.Merch, error) {
	ret := _m.Called(ctx, itemName)

	if len(ret) == 0 {
		panic("no return value specified for GetMerchByName")
	}

	var r0 entity.Merch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Merch, error)); ok {
		return rf(ctx, itemName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Merch); ok {
		r0 = rf(ctx, itemName)
	} else {
		r0 = ret.Get(0).(entity.Merch)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, itemName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetMerchByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMerchByName'
type Storage_GetMerchByName_Call struct {
	*mock.Call
}

// GetMerchByName is a helper method to define mock.On call
//   - ctx context.Context
//   - itemName string
func (_e *Storage_Expecter) GetMerchByName(ctx interface{}, itemName interface{}) *Storage_GetMerchByName_Call {
	return &Storage_GetMerchByName_Call{Call: _e.mock.On("GetMerchByName", ctx, itemName)}
}

func (_c *Storage_GetMerchByName_Call) Run(run func(ctx context.Context, itemName string)) *Storage_GetMerchByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storage_GetMerchByName_Call) Return(_a0 entity.Merch, _a1 error) *Storage_GetMerchByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_GetMerchByName_Call) RunAndReturn(run func(context.Context, string) (entity.Merch, error)) *Storage_GetMerchByName_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
