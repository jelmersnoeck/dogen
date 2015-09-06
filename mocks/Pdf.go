package mocks

import "github.com/jelmersnoeck/noscito/pdf"
import "github.com/stretchr/testify/mock"

import "bytes"

type MPdf struct {
	mock.Mock
}

func (_m *MPdf) HttpImage(url string, x float64, y float64, w float64, h float64, flow bool, tp string, link int, linkStr string) {
	_m.Called(url, x, y, w, h, flow, tp, link, linkStr)
}
func (_m *MPdf) ParseBlocks(blocks []pdf.Block) {
	_m.Called(blocks)
}
func (_m *MPdf) Bytes(buffer *bytes.Buffer) []byte {
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
