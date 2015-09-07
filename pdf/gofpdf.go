package pdf

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

type GoFpdf struct {
	fpdf   *gofpdf.Fpdf
	Layout Layout
}

// NewGoFpdf generates a new layout based on a Layout interface and uses the
// jung-kurt/gofpdf library to process all the information.
func NewGoFpdf(l Layout) (pdf *GoFpdf) {
	init := &gofpdf.InitType{
		OrientationStr: l.Orientation(),
		UnitStr:        l.Unit(),
		Size:           gofpdf.SizeType{Wd: l.Width(), Ht: l.Height()},
	}
	fpdf := gofpdf.NewCustom(init)
	fpdf.AddPage()

	pdf = new(GoFpdf)
	pdf.fpdf = fpdf
	pdf.Layout = l
	return
}

// HttpImage takes a url and registers this URL on the PDF which will then be
// used to position the image on the page and draw it. This is done in a
// goroutine so that we can download all images asynchroniously and put them on
// the page in one time.
func (f *GoFpdf) HttpImage(url string, x, y, w, h float64, flow bool, tp string, link int, linkStr string) {
	httpimg.Register(f.fpdf, url, "")
	f.fpdf.Image(url, x, y, w, h, flow, tp, link, linkStr)
}

// ParseBlocks goes through a set of blocks and parses them accordingly. The
// blocks themselve will call back to the PDF to draw the actual elements on the
// page.
func (f *GoFpdf) ParseBlocks(blocks []Block) {
	for _, block := range blocks {
		block.Parse(f)
	}
}

// Position returns the current position of the PDF.
func (f *GoFpdf) Position() (float64, float64) {
	return f.fpdf.GetXY()
}

// Position sets the position based on the two given points.
func (f *GoFpdf) SetPosition(x, y float64) {
	f.fpdf.SetXY(x, y)
}

// Text adds text to the PDF.
func (f *GoFpdf) Text(text string) {
	f.fpdf.SetFont("Times", "", 12)
	f.fpdf.MultiCell(0, 5, text, "", "", false)
}

// Bytes sends back a buffer of bytes which can be used to stream to a webpage.
func (f *GoFpdf) Bytes(buffer *bytes.Buffer) []byte {
	err := f.fpdf.Output(buffer)

	if err != nil {
		f.fpdf.SetError(err)
		return nil
	}

	return buffer.Bytes()
}
