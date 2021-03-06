// Code generated by mockery v2.5.1. DO NOT EDIT.

package prompter

import mock "github.com/stretchr/testify/mock"

// MockPrompter is an autogenerated mock type for the Prompter type
type MockPrompter struct {
	mock.Mock
}

// Choose provides a mock function with given fields: _a0, _a1
func (_m *MockPrompter) Choose(_a0 string, _a1 ...string) string {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ChooseWithDefault provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockPrompter) ChooseWithDefault(_a0 string, _a1 string, _a2 ...string) (string, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, ...string) string); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, ...string) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Confirm provides a mock function with given fields: _a0
func (_m *MockPrompter) Confirm(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Password provides a mock function with given fields: _a0
func (_m *MockPrompter) Password(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// String provides a mock function with given fields: _a0, _a1
func (_m *MockPrompter) String(_a0 string, _a1 string) string {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StringRequired provides a mock function with given fields: _a0
func (_m *MockPrompter) StringRequired(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
