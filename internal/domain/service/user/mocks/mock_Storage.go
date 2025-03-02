// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/rshelekhov/merch-store/internal/domain/entity"
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

// CreateUser provides a mock function with given fields: ctx, _a1
func (_m *Storage) CreateUser(ctx context.Context, _a1 entity.User) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type Storage_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 entity.User
func (_e *Storage_Expecter) CreateUser(ctx interface{}, _a1 interface{}) *Storage_CreateUser_Call {
	return &Storage_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, _a1)}
}

func (_c *Storage_CreateUser_Call) Run(run func(ctx context.Context, _a1 entity.User)) *Storage_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.User))
	})
	return _c
}

func (_c *Storage_CreateUser_Call) Return(_a0 error) *Storage_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_CreateUser_Call) RunAndReturn(run func(context.Context, entity.User) error) *Storage_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByName provides a mock function with given fields: ctx, username
func (_m *Storage) GetUserByName(ctx context.Context, username string) (entity.User, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByName")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetUserByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByName'
type Storage_GetUserByName_Call struct {
	*mock.Call
}

// GetUserByName is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *Storage_Expecter) GetUserByName(ctx interface{}, username interface{}) *Storage_GetUserByName_Call {
	return &Storage_GetUserByName_Call{Call: _e.mock.On("GetUserByName", ctx, username)}
}

func (_c *Storage_GetUserByName_Call) Run(run func(ctx context.Context, username string)) *Storage_GetUserByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storage_GetUserByName_Call) Return(_a0 entity.User, _a1 error) *Storage_GetUserByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_GetUserByName_Call) RunAndReturn(run func(context.Context, string) (entity.User, error)) *Storage_GetUserByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserInfoByID provides a mock function with given fields: ctx, userID
func (_m *Storage) GetUserInfoByID(ctx context.Context, userID string) (entity.UserInfo, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserInfoByID")
	}

	var r0 entity.UserInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.UserInfo, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.UserInfo); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(entity.UserInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetUserInfoByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserInfoByID'
type Storage_GetUserInfoByID_Call struct {
	*mock.Call
}

// GetUserInfoByID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *Storage_Expecter) GetUserInfoByID(ctx interface{}, userID interface{}) *Storage_GetUserInfoByID_Call {
	return &Storage_GetUserInfoByID_Call{Call: _e.mock.On("GetUserInfoByID", ctx, userID)}
}

func (_c *Storage_GetUserInfoByID_Call) Run(run func(ctx context.Context, userID string)) *Storage_GetUserInfoByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storage_GetUserInfoByID_Call) Return(_a0 entity.UserInfo, _a1 error) *Storage_GetUserInfoByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_GetUserInfoByID_Call) RunAndReturn(run func(context.Context, string) (entity.UserInfo, error)) *Storage_GetUserInfoByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserInfoByUsername provides a mock function with given fields: ctx, username
func (_m *Storage) GetUserInfoByUsername(ctx context.Context, username string) (entity.UserInfo, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetUserInfoByUsername")
	}

	var r0 entity.UserInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.UserInfo, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.UserInfo); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(entity.UserInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetUserInfoByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserInfoByUsername'
type Storage_GetUserInfoByUsername_Call struct {
	*mock.Call
}

// GetUserInfoByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *Storage_Expecter) GetUserInfoByUsername(ctx interface{}, username interface{}) *Storage_GetUserInfoByUsername_Call {
	return &Storage_GetUserInfoByUsername_Call{Call: _e.mock.On("GetUserInfoByUsername", ctx, username)}
}

func (_c *Storage_GetUserInfoByUsername_Call) Run(run func(ctx context.Context, username string)) *Storage_GetUserInfoByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storage_GetUserInfoByUsername_Call) Return(_a0 entity.UserInfo, _a1 error) *Storage_GetUserInfoByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_GetUserInfoByUsername_Call) RunAndReturn(run func(context.Context, string) (entity.UserInfo, error)) *Storage_GetUserInfoByUsername_Call {
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
