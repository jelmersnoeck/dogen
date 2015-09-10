package pdf

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
)

type GoFpdf struct {
	fpdf   *gofpdf.Fpdf
	layout Layout
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
	fpdf.SetFont("Times", "", 12)

	pdf = new(GoFpdf)
	pdf.fpdf = fpdf
	pdf.layout = l
	return
}

func (f GoFpdf) Layout() Layout {
	return f.layout
}

func (f *GoFpdf) Document() Document {
	return f.fpdf
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
