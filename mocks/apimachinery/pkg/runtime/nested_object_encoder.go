// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// NestedObjectEncoder is an autogenerated mock type for the NestedObjectEncoder type
type NestedObjectEncoder struct {
	mock.Mock
}

// EncodeNestedObjects provides a mock function with given fields: e
func (_m *NestedObjectEncoder) EncodeNestedObjects(e runtime.Encoder) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(runtime.Encoder) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
