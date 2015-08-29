package easypdf

import (
	"github.com/jung-kurt/gofpdf"
)

type Loop struct {
	Blocks []Block
}

func (l *Loop) Parse(pdf *gofpdf.Fpdf, input map[string]interface{}) {
}
