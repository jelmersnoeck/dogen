package mocks

import "github.com/stretchr/testify/mock"

import "io"

type PdfDrawer struct {
	mock.Mock
}

func (_m *PdfDrawer) AddPage() {
	_m.Called()
}
func (_m *PdfDrawer) SetError(err error) {
	_m.Called(err)
}
func (_m *PdfDrawer) Output(w io.Writer) error {
	ret := _m.Called(w)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *PdfDrawer) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
