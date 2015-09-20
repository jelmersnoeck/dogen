package mocks

import "github.com/stretchr/testify/mock"

import "io"
import "github.com/jung-kurt/gofpdf"

type Document struct {
	mock.Mock
}

func (_m *Document) AddFont(familyStr string, styleStr string, fileStr string) {
	_m.Called(familyStr, styleStr, fileStr)
}
func (_m *Document) CellFormat(w float64, h float64, txtStr string, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string) {
	_m.Called(w, h, txtStr, borderStr, ln, alignStr, fill, link, linkStr)
}
func (_m *Document) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
func (_m *Document) GetFillColor() (int, int, int) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func() int); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
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
func (_m *Document) GetMargins() (float64, float64, float64, float64) {
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

	var r2 float64
	if rf, ok := ret.Get(2).(func() float64); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(float64)
	}

	var r3 float64
	if rf, ok := ret.Get(3).(func() float64); ok {
		r3 = rf()
	} else {
		r3 = ret.Get(3).(float64)
	}

	return r0, r1, r2, r3
}
func (_m *Document) GetPageSize() (float64, float64) {
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
func (_m *Document) HTMLBasicNew() gofpdf.HTMLBasicType {
	ret := _m.Called()

	var r0 gofpdf.HTMLBasicType
	if rf, ok := ret.Get(0).(func() gofpdf.HTMLBasicType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(gofpdf.HTMLBasicType)
	}

	return r0
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
func (_m *Document) Line(x1 float64, y1 float64, x2 float64, y2 float64) {
	_m.Called(x1, y1, x2, y2)
}
func (_m *Document) LineTo(x float64, y float64) {
	_m.Called(x, y)
}
func (_m *Document) MultiCell(w float64, h float64, txtStr string, borderStr string, alignStr string, fill bool) {
	_m.Called(w, h, txtStr, borderStr, alignStr, fill)
}
func (_m *Document) PointConvert(pt float64) float64 {
	ret := _m.Called(pt)

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64) float64); ok {
		r0 = rf(pt)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}
func (_m *Document) Rect(x float64, y float64, w float64, h float64, styleStr string) {
	_m.Called(x, y, w, h, styleStr)
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
func (_m *Document) SetCellMargin(margin float64) {
	_m.Called(margin)
}
func (_m *Document) SetDrawColor(r int, g int, b int) {
	_m.Called(r, g, b)
}
func (_m *Document) SetError(err error) {
	_m.Called(err)
}
func (_m *Document) SetFillColor(r int, g int, b int) {
	_m.Called(r, g, b)
}
func (_m *Document) SetFont(familyStr string, styleStr string, size float64) {
	_m.Called(familyStr, styleStr, size)
}
func (_m *Document) SetLeftMargin(margin float64) {
	_m.Called(margin)
}
func (_m *Document) SetRightMargin(margin float64) {
	_m.Called(margin)
}
func (_m *Document) SetTextColor(r int, g int, b int) {
	_m.Called(r, g, b)
}
func (_m *Document) SetTopMargin(margin float64) {
	_m.Called(margin)
}
func (_m *Document) SetXY(x float64, y float64) {
	_m.Called(x, y)
}
func (_m *Document) SetX(x float64) {
	_m.Called(x)
}
func (_m *Document) SetY(y float64) {
	_m.Called(y)
}
func (_m *Document) Write(h float64, txtStr string) {
	_m.Called(h, txtStr)
}
