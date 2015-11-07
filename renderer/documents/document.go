// Package documents represents the available documents in our system.
package documents

import (
	"io"

	"github.com/jung-kurt/gofpdf"
)

// Document represents the actual PDF implementation which is used to generate a
// PDF document and draw to the file. This will usually be an external library
// and just serves as an interface to pass along in our own methods.
type Document interface {
	AddFont(familyStr, styleStr, fileStr string)
	AddPage()
	CellFormat(w, h float64, txtStr string, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string)
	Error() error
	GetConversionRatio() float64
	GetFillColor() (int, int, int)
	GetImageInfo(imageStr string) *gofpdf.ImageInfoType
	GetMargins() (float64, float64, float64, float64)
	GetPageSize() (float64, float64)
	GetXY() (float64, float64)
	GetY() float64
	HTMLBasicNew() (html gofpdf.HTMLBasicType)
	Image(imageNameStr string, x, y, w, h float64, flow bool, tp string, link int, linkStr string)
	ImageTypeFromMime(mimeStr string) string
	Line(x1, y1, x2, y2 float64)
	LineTo(x, y float64)
	MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool)
	PointConvert(pt float64) float64
	Rect(x, y, w, h float64, styleStr string)
	RegisterImageReader(imgName, tp string, r io.Reader) *gofpdf.ImageInfoType
	SetAutoPageBreak(auto bool, margin float64)
	SetCellMargin(margin float64)
	SetDrawColor(r, g, b int)
	SetError(err error)
	SetFillColor(r, g, b int)
	SetFont(familyStr, styleStr string, size float64)
	SetLeftMargin(margin float64)
	SetMargins(left, top, right float64)
	SetRightMargin(margin float64)
	SetTextColor(r, g, b int)
	SetTopMargin(margin float64)
	SetXY(x, y float64)
	SetX(x float64)
	SetY(y float64)
	Write(h float64, txtStr string)
	TransformBegin()
	TransformEnd()
	TransformRotate(angle, x, y float64)
	GetStringWidth(s string) float64
}
