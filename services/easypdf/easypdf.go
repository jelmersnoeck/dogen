package easypdf

import (
	"bytes"
	"github.com/jelmersnoeck/noscito/services/easypdf/template"
	"github.com/jung-kurt/gofpdf"
)

func New(l *template.Layout) (f *gofpdf.Fpdf) {
	init := &gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: l.Width, Ht: l.Height},
	}
	pdf := gofpdf.NewCustom(init)
	pdf.AddPage()

	return pdf
}

func Render(pdf *gofpdf.Fpdf) (f []byte) {
	buffer := bytes.NewBufferString("")
	err := pdf.Output(buffer)

	if err != nil {
		pdf.SetError(err)
		return
	}

	return buffer.Bytes()
}
