package easypdf

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
	"sync"
)

type EasyPDF struct {
	l   *Layout
	pdf *gofpdf.Fpdf
	Wg  sync.WaitGroup
}

func New(l Layout) (e *EasyPDF) {
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

func (e *EasyPDF) RegisterBlocks(blocks []Block, user_input map[string]interface{}) {
	for _, block := range blocks {
		block.Parse(e, user_input)
	}
}

func (e *EasyPDF) Render() (f []byte) {
	e.Wg.Wait()
	buffer := bytes.NewBufferString("")
	err := e.pdf.Output(buffer)

	if err != nil {
		e.pdf.SetError(err)
		return
	}

	return buffer.Bytes()
}
