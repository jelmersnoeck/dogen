// Package PDF contains functionality to draw data onto a PDF and render said
// PDF to a file or string buffer.
package pdf

import (
	"bytes"
	"io"

	"github.com/jung-kurt/gofpdf"
)

type Pdf interface {
	Bytes(buffer *bytes.Buffer) []byte
	Layout() Layout
	Document() Document
}

type Document interface {
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

// ParseBlocks goes through a set of blocks and parses them accordingly. The
// blocks themselve will call back to the PDF to draw the actual elements on the
// page.
func ParseBlocks(pdf Pdf, blocks []Block) {
	for _, block := range blocks {
		block.Parse(pdf.Document())
	}
}
