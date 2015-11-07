package mocks

import (
	"github.com/jelmersnoeck/dogen/renderer/layouts"
	"github.com/jelmersnoeck/dogen/renderer/template"
	"github.com/stretchr/testify/mock"
)

import "sync"

type Template struct {
	mock.Mock
}

func (_m *Template) AddError(err error) {
	_m.Called(err)
}
func (_m *Template) Blocks() []template.Block {
	ret := _m.Called()

	var r0 []template.Block
	if rf, ok := ret.Get(0).(func() []template.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]template.Block)
		}
	}

	return r0
}
func (_m *Template) Errors() []error {
	ret := _m.Called()

	var r0 []error
	if rf, ok := ret.Get(0).(func() []error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}
func (_m *Template) Layout() layouts.Layout {
	ret := _m.Called()

	var r0 layouts.Layout
	if rf, ok := ret.Get(0).(func() layouts.Layout); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(layouts.Layout)
	}

	return r0
}
func (_m *Template) LoadBlock(raw_block map[string]interface{}, raw_input map[string]interface{}) template.Block {
	ret := _m.Called(raw_block, raw_input)

	var r0 template.Block
	if rf, ok := ret.Get(0).(func(map[string]interface{}, map[string]interface{}) template.Block); ok {
		r0 = rf(raw_block, raw_input)
	} else {
		r0 = ret.Get(0).(template.Block)
	}

	return r0
}
func (_m *Template) LoadBlocks(user_input interface{}) {
	_m.Called(user_input)
}
func (_m *Template) WaitGroup() *sync.WaitGroup {
	ret := _m.Called()

	var r0 *sync.WaitGroup
	if rf, ok := ret.Get(0).(func() *sync.WaitGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sync.WaitGroup)
		}
	}

	return r0
}
