package block

import (
	"github.com/jung-kurt/gofpdf"
)

type Text struct {
	Value string
}

func (t *Text) Parse(pdf *gofpdf.Fpdf, input map[string]interface{}) {
}
