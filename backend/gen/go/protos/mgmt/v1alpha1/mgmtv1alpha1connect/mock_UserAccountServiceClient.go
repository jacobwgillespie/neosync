// Code generated by mockery. DO NOT EDIT.

package mgmtv1alpha1connect

import (
	context "context"

	connect "connectrpc.com/connect"

	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MockUserAccountServiceClient is an autogenerated mock type for the UserAccountServiceClient type
type MockUserAccountServiceClient struct {
	mock.Mock
}

type MockUserAccountServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserAccountServiceClient) EXPECT() *MockUserAccountServiceClient_Expecter {
	return &MockUserAccountServiceClient_Expecter{mock: &_m.Mock}
}

// ConvertPersonalToTeamAccount provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) ConvertPersonalToTeamAccount(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.ConvertPersonalToTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.ConvertPersonalToTeamAccountResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.ConvertPersonalToTeamAccountResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.ConvertPersonalToTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.ConvertPersonalToTeamAccountResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.ConvertPersonalToTeamAccountRequest]) *connect.Response[mgmtv1alpha1.ConvertPersonalToTeamAccountResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.ConvertPersonalToTeamAccountResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.ConvertPersonalToTeamAccountRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConvertPersonalToTeamAccount'
type MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call struct {
	*mock.Call
}

// ConvertPersonalToTeamAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.ConvertPersonalToTeamAccountRequest]
func (_e *MockUserAccountServiceClient_Expecter) ConvertPersonalToTeamAccount(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call {
	return &MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call{Call: _e.mock.On("ConvertPersonalToTeamAccount", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.ConvertPersonalToTeamAccountRequest])) *MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.ConvertPersonalToTeamAccountRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call) Return(_a0 *connect.Response[mgmtv1alpha1.ConvertPersonalToTeamAccountResponse], _a1 error) *MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.ConvertPersonalToTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.ConvertPersonalToTeamAccountResponse], error)) *MockUserAccountServiceClient_ConvertPersonalToTeamAccount_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTeamAccount provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) CreateTeamAccount(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CreateTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.CreateTeamAccountResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.CreateTeamAccountResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.CreateTeamAccountResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateTeamAccountRequest]) *connect.Response[mgmtv1alpha1.CreateTeamAccountResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.CreateTeamAccountResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateTeamAccountRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_CreateTeamAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTeamAccount'
type MockUserAccountServiceClient_CreateTeamAccount_Call struct {
	*mock.Call
}

// CreateTeamAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.CreateTeamAccountRequest]
func (_e *MockUserAccountServiceClient_Expecter) CreateTeamAccount(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_CreateTeamAccount_Call {
	return &MockUserAccountServiceClient_CreateTeamAccount_Call{Call: _e.mock.On("CreateTeamAccount", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_CreateTeamAccount_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CreateTeamAccountRequest])) *MockUserAccountServiceClient_CreateTeamAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.CreateTeamAccountRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_CreateTeamAccount_Call) Return(_a0 *connect.Response[mgmtv1alpha1.CreateTeamAccountResponse], _a1 error) *MockUserAccountServiceClient_CreateTeamAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_CreateTeamAccount_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.CreateTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.CreateTeamAccountResponse], error)) *MockUserAccountServiceClient_CreateTeamAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountTemporalConfig provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) GetAccountTemporalConfig(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetAccountTemporalConfigRequest]) (*connect.Response[mgmtv1alpha1.GetAccountTemporalConfigResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.GetAccountTemporalConfigResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAccountTemporalConfigRequest]) (*connect.Response[mgmtv1alpha1.GetAccountTemporalConfigResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAccountTemporalConfigRequest]) *connect.Response[mgmtv1alpha1.GetAccountTemporalConfigResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetAccountTemporalConfigResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetAccountTemporalConfigRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_GetAccountTemporalConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountTemporalConfig'
type MockUserAccountServiceClient_GetAccountTemporalConfig_Call struct {
	*mock.Call
}

// GetAccountTemporalConfig is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetAccountTemporalConfigRequest]
func (_e *MockUserAccountServiceClient_Expecter) GetAccountTemporalConfig(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_GetAccountTemporalConfig_Call {
	return &MockUserAccountServiceClient_GetAccountTemporalConfig_Call{Call: _e.mock.On("GetAccountTemporalConfig", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_GetAccountTemporalConfig_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetAccountTemporalConfigRequest])) *MockUserAccountServiceClient_GetAccountTemporalConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetAccountTemporalConfigRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_GetAccountTemporalConfig_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetAccountTemporalConfigResponse], _a1 error) *MockUserAccountServiceClient_GetAccountTemporalConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_GetAccountTemporalConfig_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetAccountTemporalConfigRequest]) (*connect.Response[mgmtv1alpha1.GetAccountTemporalConfigResponse], error)) *MockUserAccountServiceClient_GetAccountTemporalConfig_Call {
	_c.Call.Return(run)
	return _c
}

// GetTeamAccountInvites provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) GetTeamAccountInvites(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetTeamAccountInvitesRequest]) (*connect.Response[mgmtv1alpha1.GetTeamAccountInvitesResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.GetTeamAccountInvitesResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetTeamAccountInvitesRequest]) (*connect.Response[mgmtv1alpha1.GetTeamAccountInvitesResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetTeamAccountInvitesRequest]) *connect.Response[mgmtv1alpha1.GetTeamAccountInvitesResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetTeamAccountInvitesResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetTeamAccountInvitesRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_GetTeamAccountInvites_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTeamAccountInvites'
type MockUserAccountServiceClient_GetTeamAccountInvites_Call struct {
	*mock.Call
}

// GetTeamAccountInvites is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetTeamAccountInvitesRequest]
func (_e *MockUserAccountServiceClient_Expecter) GetTeamAccountInvites(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_GetTeamAccountInvites_Call {
	return &MockUserAccountServiceClient_GetTeamAccountInvites_Call{Call: _e.mock.On("GetTeamAccountInvites", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_GetTeamAccountInvites_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetTeamAccountInvitesRequest])) *MockUserAccountServiceClient_GetTeamAccountInvites_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetTeamAccountInvitesRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_GetTeamAccountInvites_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetTeamAccountInvitesResponse], _a1 error) *MockUserAccountServiceClient_GetTeamAccountInvites_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_GetTeamAccountInvites_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetTeamAccountInvitesRequest]) (*connect.Response[mgmtv1alpha1.GetTeamAccountInvitesResponse], error)) *MockUserAccountServiceClient_GetTeamAccountInvites_Call {
	_c.Call.Return(run)
	return _c
}

// GetTeamAccountMembers provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) GetTeamAccountMembers(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetTeamAccountMembersRequest]) (*connect.Response[mgmtv1alpha1.GetTeamAccountMembersResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.GetTeamAccountMembersResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetTeamAccountMembersRequest]) (*connect.Response[mgmtv1alpha1.GetTeamAccountMembersResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetTeamAccountMembersRequest]) *connect.Response[mgmtv1alpha1.GetTeamAccountMembersResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetTeamAccountMembersResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetTeamAccountMembersRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_GetTeamAccountMembers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTeamAccountMembers'
type MockUserAccountServiceClient_GetTeamAccountMembers_Call struct {
	*mock.Call
}

// GetTeamAccountMembers is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetTeamAccountMembersRequest]
func (_e *MockUserAccountServiceClient_Expecter) GetTeamAccountMembers(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_GetTeamAccountMembers_Call {
	return &MockUserAccountServiceClient_GetTeamAccountMembers_Call{Call: _e.mock.On("GetTeamAccountMembers", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_GetTeamAccountMembers_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetTeamAccountMembersRequest])) *MockUserAccountServiceClient_GetTeamAccountMembers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetTeamAccountMembersRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_GetTeamAccountMembers_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetTeamAccountMembersResponse], _a1 error) *MockUserAccountServiceClient_GetTeamAccountMembers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_GetTeamAccountMembers_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetTeamAccountMembersRequest]) (*connect.Response[mgmtv1alpha1.GetTeamAccountMembersResponse], error)) *MockUserAccountServiceClient_GetTeamAccountMembers_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) GetUser(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetUserRequest]) (*connect.Response[mgmtv1alpha1.GetUserResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.GetUserResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetUserRequest]) (*connect.Response[mgmtv1alpha1.GetUserResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetUserRequest]) *connect.Response[mgmtv1alpha1.GetUserResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetUserResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetUserRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type MockUserAccountServiceClient_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetUserRequest]
func (_e *MockUserAccountServiceClient_Expecter) GetUser(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_GetUser_Call {
	return &MockUserAccountServiceClient_GetUser_Call{Call: _e.mock.On("GetUser", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_GetUser_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetUserRequest])) *MockUserAccountServiceClient_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetUserRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_GetUser_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetUserResponse], _a1 error) *MockUserAccountServiceClient_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_GetUser_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetUserRequest]) (*connect.Response[mgmtv1alpha1.GetUserResponse], error)) *MockUserAccountServiceClient_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserAccounts provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) GetUserAccounts(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetUserAccountsRequest]) (*connect.Response[mgmtv1alpha1.GetUserAccountsResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.GetUserAccountsResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetUserAccountsRequest]) (*connect.Response[mgmtv1alpha1.GetUserAccountsResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetUserAccountsRequest]) *connect.Response[mgmtv1alpha1.GetUserAccountsResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetUserAccountsResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetUserAccountsRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_GetUserAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserAccounts'
type MockUserAccountServiceClient_GetUserAccounts_Call struct {
	*mock.Call
}

// GetUserAccounts is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetUserAccountsRequest]
func (_e *MockUserAccountServiceClient_Expecter) GetUserAccounts(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_GetUserAccounts_Call {
	return &MockUserAccountServiceClient_GetUserAccounts_Call{Call: _e.mock.On("GetUserAccounts", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_GetUserAccounts_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetUserAccountsRequest])) *MockUserAccountServiceClient_GetUserAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetUserAccountsRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_GetUserAccounts_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetUserAccountsResponse], _a1 error) *MockUserAccountServiceClient_GetUserAccounts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_GetUserAccounts_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetUserAccountsRequest]) (*connect.Response[mgmtv1alpha1.GetUserAccountsResponse], error)) *MockUserAccountServiceClient_GetUserAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// InviteUserToTeamAccount provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) InviteUserToTeamAccount(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.InviteUserToTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.InviteUserToTeamAccountResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.InviteUserToTeamAccountResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.InviteUserToTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.InviteUserToTeamAccountResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.InviteUserToTeamAccountRequest]) *connect.Response[mgmtv1alpha1.InviteUserToTeamAccountResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.InviteUserToTeamAccountResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.InviteUserToTeamAccountRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_InviteUserToTeamAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InviteUserToTeamAccount'
type MockUserAccountServiceClient_InviteUserToTeamAccount_Call struct {
	*mock.Call
}

// InviteUserToTeamAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.InviteUserToTeamAccountRequest]
func (_e *MockUserAccountServiceClient_Expecter) InviteUserToTeamAccount(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_InviteUserToTeamAccount_Call {
	return &MockUserAccountServiceClient_InviteUserToTeamAccount_Call{Call: _e.mock.On("InviteUserToTeamAccount", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_InviteUserToTeamAccount_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.InviteUserToTeamAccountRequest])) *MockUserAccountServiceClient_InviteUserToTeamAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.InviteUserToTeamAccountRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_InviteUserToTeamAccount_Call) Return(_a0 *connect.Response[mgmtv1alpha1.InviteUserToTeamAccountResponse], _a1 error) *MockUserAccountServiceClient_InviteUserToTeamAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_InviteUserToTeamAccount_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.InviteUserToTeamAccountRequest]) (*connect.Response[mgmtv1alpha1.InviteUserToTeamAccountResponse], error)) *MockUserAccountServiceClient_InviteUserToTeamAccount_Call {
	_c.Call.Return(run)
	return _c
}

// IsUserInAccount provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) IsUserInAccount(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.IsUserInAccountRequest]) (*connect.Response[mgmtv1alpha1.IsUserInAccountResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.IsUserInAccountResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.IsUserInAccountRequest]) (*connect.Response[mgmtv1alpha1.IsUserInAccountResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.IsUserInAccountRequest]) *connect.Response[mgmtv1alpha1.IsUserInAccountResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.IsUserInAccountResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.IsUserInAccountRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_IsUserInAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUserInAccount'
type MockUserAccountServiceClient_IsUserInAccount_Call struct {
	*mock.Call
}

// IsUserInAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.IsUserInAccountRequest]
func (_e *MockUserAccountServiceClient_Expecter) IsUserInAccount(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_IsUserInAccount_Call {
	return &MockUserAccountServiceClient_IsUserInAccount_Call{Call: _e.mock.On("IsUserInAccount", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_IsUserInAccount_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.IsUserInAccountRequest])) *MockUserAccountServiceClient_IsUserInAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.IsUserInAccountRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_IsUserInAccount_Call) Return(_a0 *connect.Response[mgmtv1alpha1.IsUserInAccountResponse], _a1 error) *MockUserAccountServiceClient_IsUserInAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_IsUserInAccount_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.IsUserInAccountRequest]) (*connect.Response[mgmtv1alpha1.IsUserInAccountResponse], error)) *MockUserAccountServiceClient_IsUserInAccount_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveTeamAccountInvite provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) RemoveTeamAccountInvite(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.RemoveTeamAccountInviteRequest]) (*connect.Response[mgmtv1alpha1.RemoveTeamAccountInviteResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.RemoveTeamAccountInviteResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.RemoveTeamAccountInviteRequest]) (*connect.Response[mgmtv1alpha1.RemoveTeamAccountInviteResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.RemoveTeamAccountInviteRequest]) *connect.Response[mgmtv1alpha1.RemoveTeamAccountInviteResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.RemoveTeamAccountInviteResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.RemoveTeamAccountInviteRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_RemoveTeamAccountInvite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveTeamAccountInvite'
type MockUserAccountServiceClient_RemoveTeamAccountInvite_Call struct {
	*mock.Call
}

// RemoveTeamAccountInvite is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.RemoveTeamAccountInviteRequest]
func (_e *MockUserAccountServiceClient_Expecter) RemoveTeamAccountInvite(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_RemoveTeamAccountInvite_Call {
	return &MockUserAccountServiceClient_RemoveTeamAccountInvite_Call{Call: _e.mock.On("RemoveTeamAccountInvite", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_RemoveTeamAccountInvite_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.RemoveTeamAccountInviteRequest])) *MockUserAccountServiceClient_RemoveTeamAccountInvite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.RemoveTeamAccountInviteRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_RemoveTeamAccountInvite_Call) Return(_a0 *connect.Response[mgmtv1alpha1.RemoveTeamAccountInviteResponse], _a1 error) *MockUserAccountServiceClient_RemoveTeamAccountInvite_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_RemoveTeamAccountInvite_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.RemoveTeamAccountInviteRequest]) (*connect.Response[mgmtv1alpha1.RemoveTeamAccountInviteResponse], error)) *MockUserAccountServiceClient_RemoveTeamAccountInvite_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveTeamAccountMember provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) RemoveTeamAccountMember(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.RemoveTeamAccountMemberRequest]) (*connect.Response[mgmtv1alpha1.RemoveTeamAccountMemberResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.RemoveTeamAccountMemberResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.RemoveTeamAccountMemberRequest]) (*connect.Response[mgmtv1alpha1.RemoveTeamAccountMemberResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.RemoveTeamAccountMemberRequest]) *connect.Response[mgmtv1alpha1.RemoveTeamAccountMemberResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.RemoveTeamAccountMemberResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.RemoveTeamAccountMemberRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_RemoveTeamAccountMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveTeamAccountMember'
type MockUserAccountServiceClient_RemoveTeamAccountMember_Call struct {
	*mock.Call
}

// RemoveTeamAccountMember is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.RemoveTeamAccountMemberRequest]
func (_e *MockUserAccountServiceClient_Expecter) RemoveTeamAccountMember(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_RemoveTeamAccountMember_Call {
	return &MockUserAccountServiceClient_RemoveTeamAccountMember_Call{Call: _e.mock.On("RemoveTeamAccountMember", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_RemoveTeamAccountMember_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.RemoveTeamAccountMemberRequest])) *MockUserAccountServiceClient_RemoveTeamAccountMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.RemoveTeamAccountMemberRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_RemoveTeamAccountMember_Call) Return(_a0 *connect.Response[mgmtv1alpha1.RemoveTeamAccountMemberResponse], _a1 error) *MockUserAccountServiceClient_RemoveTeamAccountMember_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_RemoveTeamAccountMember_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.RemoveTeamAccountMemberRequest]) (*connect.Response[mgmtv1alpha1.RemoveTeamAccountMemberResponse], error)) *MockUserAccountServiceClient_RemoveTeamAccountMember_Call {
	_c.Call.Return(run)
	return _c
}

// SetAccountTemporalConfig provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) SetAccountTemporalConfig(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.SetAccountTemporalConfigRequest]) (*connect.Response[mgmtv1alpha1.SetAccountTemporalConfigResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.SetAccountTemporalConfigResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.SetAccountTemporalConfigRequest]) (*connect.Response[mgmtv1alpha1.SetAccountTemporalConfigResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.SetAccountTemporalConfigRequest]) *connect.Response[mgmtv1alpha1.SetAccountTemporalConfigResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.SetAccountTemporalConfigResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.SetAccountTemporalConfigRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_SetAccountTemporalConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAccountTemporalConfig'
type MockUserAccountServiceClient_SetAccountTemporalConfig_Call struct {
	*mock.Call
}

// SetAccountTemporalConfig is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.SetAccountTemporalConfigRequest]
func (_e *MockUserAccountServiceClient_Expecter) SetAccountTemporalConfig(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_SetAccountTemporalConfig_Call {
	return &MockUserAccountServiceClient_SetAccountTemporalConfig_Call{Call: _e.mock.On("SetAccountTemporalConfig", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_SetAccountTemporalConfig_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.SetAccountTemporalConfigRequest])) *MockUserAccountServiceClient_SetAccountTemporalConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.SetAccountTemporalConfigRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_SetAccountTemporalConfig_Call) Return(_a0 *connect.Response[mgmtv1alpha1.SetAccountTemporalConfigResponse], _a1 error) *MockUserAccountServiceClient_SetAccountTemporalConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_SetAccountTemporalConfig_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.SetAccountTemporalConfigRequest]) (*connect.Response[mgmtv1alpha1.SetAccountTemporalConfigResponse], error)) *MockUserAccountServiceClient_SetAccountTemporalConfig_Call {
	_c.Call.Return(run)
	return _c
}

// SetPersonalAccount provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) SetPersonalAccount(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.SetPersonalAccountRequest]) (*connect.Response[mgmtv1alpha1.SetPersonalAccountResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.SetPersonalAccountResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.SetPersonalAccountRequest]) (*connect.Response[mgmtv1alpha1.SetPersonalAccountResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.SetPersonalAccountRequest]) *connect.Response[mgmtv1alpha1.SetPersonalAccountResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.SetPersonalAccountResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.SetPersonalAccountRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_SetPersonalAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPersonalAccount'
type MockUserAccountServiceClient_SetPersonalAccount_Call struct {
	*mock.Call
}

// SetPersonalAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.SetPersonalAccountRequest]
func (_e *MockUserAccountServiceClient_Expecter) SetPersonalAccount(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_SetPersonalAccount_Call {
	return &MockUserAccountServiceClient_SetPersonalAccount_Call{Call: _e.mock.On("SetPersonalAccount", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_SetPersonalAccount_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.SetPersonalAccountRequest])) *MockUserAccountServiceClient_SetPersonalAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.SetPersonalAccountRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_SetPersonalAccount_Call) Return(_a0 *connect.Response[mgmtv1alpha1.SetPersonalAccountResponse], _a1 error) *MockUserAccountServiceClient_SetPersonalAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_SetPersonalAccount_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.SetPersonalAccountRequest]) (*connect.Response[mgmtv1alpha1.SetPersonalAccountResponse], error)) *MockUserAccountServiceClient_SetPersonalAccount_Call {
	_c.Call.Return(run)
	return _c
}

// SetUser provides a mock function with given fields: _a0, _a1
func (_m *MockUserAccountServiceClient) SetUser(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.SetUserRequest]) (*connect.Response[mgmtv1alpha1.SetUserResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.SetUserResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.SetUserRequest]) (*connect.Response[mgmtv1alpha1.SetUserResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.SetUserRequest]) *connect.Response[mgmtv1alpha1.SetUserResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.SetUserResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.SetUserRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserAccountServiceClient_SetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUser'
type MockUserAccountServiceClient_SetUser_Call struct {
	*mock.Call
}

// SetUser is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.SetUserRequest]
func (_e *MockUserAccountServiceClient_Expecter) SetUser(_a0 interface{}, _a1 interface{}) *MockUserAccountServiceClient_SetUser_Call {
	return &MockUserAccountServiceClient_SetUser_Call{Call: _e.mock.On("SetUser", _a0, _a1)}
}

func (_c *MockUserAccountServiceClient_SetUser_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.SetUserRequest])) *MockUserAccountServiceClient_SetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.SetUserRequest]))
	})
	return _c
}

func (_c *MockUserAccountServiceClient_SetUser_Call) Return(_a0 *connect.Response[mgmtv1alpha1.SetUserResponse], _a1 error) *MockUserAccountServiceClient_SetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserAccountServiceClient_SetUser_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.SetUserRequest]) (*connect.Response[mgmtv1alpha1.SetUserResponse], error)) *MockUserAccountServiceClient_SetUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserAccountServiceClient creates a new instance of MockUserAccountServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserAccountServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserAccountServiceClient {
	mock := &MockUserAccountServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
