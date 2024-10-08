// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "github.com/nu1lspaxe/go-0x001/server/domain"

	mock "github.com/stretchr/testify/mock"
)

// DietRepository is an autogenerated mock type for the DietRepository type
type DietRepository struct {
	mock.Mock
}

// GetById provides a mock function with given fields: ctx, id
func (_m *DietRepository) GetById(ctx context.Context, id string) (*domain.Diet, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *domain.Diet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Diet, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Diet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Diet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, d
func (_m *DietRepository) Store(ctx context.Context, d *domain.Diet) error {
	ret := _m.Called(ctx, d)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Diet) error); ok {
		r0 = rf(ctx, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDietRepository creates a new instance of DietRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDietRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *DietRepository {
	mock := &DietRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
