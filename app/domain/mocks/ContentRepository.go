// Code generated by mockery v2.2.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/meroedu/course-api/app/domain"
	mock "github.com/stretchr/testify/mock"
)

// ContentRepository is an autogenerated mock type for the ContentRepository type
type ContentRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ContentRepository) GetByID(ctx context.Context, id int64) (domain.Content, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Content
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Content); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Content)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
