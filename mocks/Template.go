package mocks

import "github.com/jelmersnoeck/noscito/pdf"
import "github.com/stretchr/testify/mock"

type Template struct {
	mock.Mock
}

func (_m *Template) AddError(err error) {
	_m.Called(err)
}
func (_m *Template) Blocks() []pdf.Block {
	ret := _m.Called()

	var r0 []pdf.Block
	if rf, ok := ret.Get(0).(func() []pdf.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pdf.Block)
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
func (_m *Template) Layout() pdf.Layout {
	ret := _m.Called()

	var r0 pdf.Layout
	if rf, ok := ret.Get(0).(func() pdf.Layout); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pdf.Layout)
	}

	return r0
}
func (_m *Template) LoadBlock(raw_block map[string]interface{}, raw_input map[string]interface{}) pdf.Block {
	ret := _m.Called(raw_block, raw_input)

	var r0 pdf.Block
	if rf, ok := ret.Get(0).(func(map[string]interface{}, map[string]interface{}) pdf.Block); ok {
		r0 = rf(raw_block, raw_input)
	} else {
		r0 = ret.Get(0).(pdf.Block)
	}

	return r0
}
func (_m *Template) LoadBlocks(user_input interface{}) {
	_m.Called(user_input)
}
