package mocks

import "github.com/jelmersnoeck/dogen/renderer/pdf"
import "github.com/stretchr/testify/mock"

import "bytes"

type Pdf struct {
	mock.Mock
}

func (_m *Pdf) Bytes(buffer *bytes.Buffer) []byte {
	ret := _m.Called(buffer)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*bytes.Buffer) []byte); ok {
		r0 = rf(buffer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}
func (_m *Pdf) LoadFonts(name string, styles map[string]string) {
	_m.Called(name, styles)
}
func (_m *Pdf) Layout() pdf.Layout {
	ret := _m.Called()

	var r0 pdf.Layout
	if rf, ok := ret.Get(0).(func() pdf.Layout); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pdf.Layout)
	}

	return r0
}
func (_m *Pdf) Document() pdf.Document {
	ret := _m.Called()

	var r0 pdf.Document
	if rf, ok := ret.Get(0).(func() pdf.Document); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pdf.Document)
	}

	return r0
}
