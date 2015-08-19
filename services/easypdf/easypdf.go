package easypdf

import (
	"bytes"
	"github.com/jelmersnoeck/noscito/services/easypdf/block"
	"github.com/jelmersnoeck/noscito/services/easypdf/template"
	"github.com/jung-kurt/gofpdf"
)

func New(l *template.Layout) (f *gofpdf.Fpdf) {
	init := &gofpdf.InitType{
		OrientationStr: l.Orientation,
		UnitStr:        l.Unit,
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

func LoadBlocks(pdf *gofpdf.Fpdf, blocks []block.BlockItem, data map[string]interface{}) {
	for _, input_block := range blocks {
		input_block.Parse(pdf, data)
	}
}
