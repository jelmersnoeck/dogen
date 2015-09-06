package mocks

import "github.com/jelmersnoeck/noscito/pdf"
import "github.com/stretchr/testify/mock"

type MTemplate struct {
	mock.Mock
}

func (_m *MTemplate) Layout() pdf.Layout {
	ret := _m.Called()

	var r0 pdf.Layout
	if rf, ok := ret.Get(0).(func() pdf.Layout); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pdf.Layout)
	}

	return r0
}
func (_m *MTemplate) LoadBlock(raw_block map[string]interface{}, raw_input map[string]interface{}) {
	_m.Called(raw_block, raw_input)
}
func (_m *MTemplate) LoadBlocks(user_input interface{}) {
	_m.Called(user_input)
}
func (_m *MTemplate) Blocks() []pdf.Block {
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
