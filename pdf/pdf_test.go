package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/assert"
)

func TestNewPdf(t *testing.T) {
	easypdf := pdf.NewPdf()

	var ep *pdf.EasyPdf
	var gp *pdf.PdfDrawer

	assert.IsType(t, ep, easypdf)
	assert.IsType(t, gp, easypdf.Drawer())
}
