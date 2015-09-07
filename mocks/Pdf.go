package mocks

import "github.com/stretchr/testify/mock"

import "bytes"

type Pdf struct {
	mock.Mock
}

func (_m *Pdf) HttpImage(url string, x float64, y float64, w float64, h float64, flow bool, tp string, link int, linkStr string) {
	_m.Called(url, x, y, w, h, flow, tp, link, linkStr)
}
func (_m *Pdf) Position() (float64, float64) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}
func (_m *Pdf) SetPosition(x float64, y float64) {
	_m.Called(x, y)
}
func (_m *Pdf) Text(text string) {
	_m.Called(text)
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
