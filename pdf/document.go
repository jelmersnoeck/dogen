package pdf

import (
	"io"

	"github.com/jung-kurt/gofpdf"
)

// Document represents the actual PDF implementation which is used to generate a
// PDF document and draw to the file. This will usually be an external library
// and just serves as an interface to pass along in our own methods.
type Document interface {
	Error() error
	GetConversionRatio() float64
	GetImageInfo(imageStr string) *gofpdf.ImageInfoType
	GetXY() (float64, float64)
	Image(imageNameStr string, x, y, w, h float64, flow bool, tp string, link int, linkStr string)
	ImageTypeFromMime(mimeStr string) string
	MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool)
	RegisterImageReader(imgName, tp string, r io.Reader) *gofpdf.ImageInfoType
	SetXY(x, y float64)
	SetError(err error)
}
