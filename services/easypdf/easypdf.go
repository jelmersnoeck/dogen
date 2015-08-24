package easypdf

import (
	"bytes"
	"github.com/jelmersnoeck/noscito/services/easypdf/block"
	"github.com/jelmersnoeck/noscito/services/easypdf/template"
	"github.com/jung-kurt/gofpdf"
)

type EasyPDF struct {
	l   *template.Layout
	pdf *gofpdf.Fpdf
}

func New(l template.Layout) (e *EasyPDF) {
	init := &gofpdf.InitType{
		OrientationStr: l.Orientation,
		UnitStr:        l.Unit,
		Size:           gofpdf.SizeType{Wd: l.Width, Ht: l.Height},
	}
	pdf := gofpdf.NewCustom(init)
	pdf.AddPage()

	e = new(EasyPDF)
	e.l = &l
	e.pdf = pdf

	return e
}

func (e *EasyPDF) RegisterBlocks(blocks []block.Block, user_input map[string]interface{}) {
	for _, block := range blocks {
		block.Parse(e.pdf, user_input)
	}
}

func (e *EasyPDF) Render() (f []byte) {
	buffer := bytes.NewBufferString("")
	err := e.pdf.Output(buffer)

	if err != nil {
		e.pdf.SetError(err)
		return
	}

	return buffer.Bytes()
}
