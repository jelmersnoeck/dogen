package mocks

import "github.com/stretchr/testify/mock"

import "io"
import "github.com/jung-kurt/gofpdf"

type Document struct {
	mock.Mock
}

func (_m *Document) GetConversionRatio() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}
func (_m *Document) GetImageInfo(imageStr string) *gofpdf.ImageInfoType {
	ret := _m.Called(imageStr)

	var r0 *gofpdf.ImageInfoType
	if rf, ok := ret.Get(0).(func(string) *gofpdf.ImageInfoType); ok {
		r0 = rf(imageStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gofpdf.ImageInfoType)
		}
	}

	return r0
}
func (_m *Document) GetXY() (float64, float64) {
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
func (_m *Document) Image(imageNameStr string, x float64, y float64, w float64, h float64, flow bool, tp string, link int, linkStr string) {
	_m.Called(imageNameStr, x, y, w, h, flow, tp, link, linkStr)
}
func (_m *Document) ImageTypeFromMime(mimeStr string) string {
	ret := _m.Called(mimeStr)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(mimeStr)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
func (_m *Document) MultiCell(w float64, h float64, txtStr string, borderStr string, alignStr string, fill bool) {
	_m.Called(w, h, txtStr, borderStr, alignStr, fill)
}
func (_m *Document) RegisterImageReader(imgName string, tp string, r io.Reader) *gofpdf.ImageInfoType {
	ret := _m.Called(imgName, tp, r)

	var r0 *gofpdf.ImageInfoType
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) *gofpdf.ImageInfoType); ok {
		r0 = rf(imgName, tp, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gofpdf.ImageInfoType)
		}
	}

	return r0
}
func (_m *Document) SetXY(x float64, y float64) {
	_m.Called(x, y)
}
func (_m *Document) SetError(err error) {
	_m.Called(err)
}
