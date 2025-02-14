// Code generated by mockery. DO NOT EDIT.

package clientmanager

import (
	context "context"

	db_queries "github.com/nucleuscloud/neosync/backend/gen/go/db"
	mock "github.com/stretchr/testify/mock"

	pg_models "github.com/nucleuscloud/neosync/backend/sql/postgresql/models"

	pgtype "github.com/jackc/pgx/v5/pgtype"
)

// MockDB is an autogenerated mock type for the DB type
type MockDB struct {
	mock.Mock
}

type MockDB_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDB) EXPECT() *MockDB_Expecter {
	return &MockDB_Expecter{mock: &_m.Mock}
}

// GetTemporalConfigByAccount provides a mock function with given fields: ctx, db, accountId
func (_m *MockDB) GetTemporalConfigByAccount(ctx context.Context, db db_queries.DBTX, accountId pgtype.UUID) (*pg_models.TemporalConfig, error) {
	ret := _m.Called(ctx, db, accountId)

	var r0 *pg_models.TemporalConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db_queries.DBTX, pgtype.UUID) (*pg_models.TemporalConfig, error)); ok {
		return rf(ctx, db, accountId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db_queries.DBTX, pgtype.UUID) *pg_models.TemporalConfig); ok {
		r0 = rf(ctx, db, accountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg_models.TemporalConfig)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db_queries.DBTX, pgtype.UUID) error); ok {
		r1 = rf(ctx, db, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDB_GetTemporalConfigByAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemporalConfigByAccount'
type MockDB_GetTemporalConfigByAccount_Call struct {
	*mock.Call
}

// GetTemporalConfigByAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - db db_queries.DBTX
//   - accountId pgtype.UUID
func (_e *MockDB_Expecter) GetTemporalConfigByAccount(ctx interface{}, db interface{}, accountId interface{}) *MockDB_GetTemporalConfigByAccount_Call {
	return &MockDB_GetTemporalConfigByAccount_Call{Call: _e.mock.On("GetTemporalConfigByAccount", ctx, db, accountId)}
}

func (_c *MockDB_GetTemporalConfigByAccount_Call) Run(run func(ctx context.Context, db db_queries.DBTX, accountId pgtype.UUID)) *MockDB_GetTemporalConfigByAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db_queries.DBTX), args[2].(pgtype.UUID))
	})
	return _c
}

func (_c *MockDB_GetTemporalConfigByAccount_Call) Return(_a0 *pg_models.TemporalConfig, _a1 error) *MockDB_GetTemporalConfigByAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDB_GetTemporalConfigByAccount_Call) RunAndReturn(run func(context.Context, db_queries.DBTX, pgtype.UUID) (*pg_models.TemporalConfig, error)) *MockDB_GetTemporalConfigByAccount_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDB creates a new instance of MockDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDB {
	mock := &MockDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
