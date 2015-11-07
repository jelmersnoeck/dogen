package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/dogen/renderer/documents/pdf"
	"github.com/jelmersnoeck/dogen/renderer/layouts"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoFpdfSuite struct {
	suite.Suite
}

func TestGoFpdfSuite(t *testing.T) {
	suite.Run(t, new(GoFpdfSuite))
}

func (s *GoFpdfSuite) TestNewGoFpdf() {
	easypdf := newPdf()

	var ep *pdf.GoFpdf
	assert.IsType(s.T(), ep, easypdf)
}

func (s *GoFpdfSuite) TestDefaultMargin() {
	easypdf := newPdf()

	left, top, right, bottom := easypdf.Document().GetMargins()
	assert.Equal(s.T(), 0.0, left)
	assert.Equal(s.T(), 0.0, top)
	assert.Equal(s.T(), 0.0, right)
	assert.Equal(s.T(), 0.0, bottom)
}

func (s *GoFpdfSuite) TestDocument() {
	easypdf := newPdf()

	var doc *gofpdf.Fpdf
	assert.IsType(s.T(), doc, easypdf.Document())
}

func (s *GoFpdfSuite) TestLayout() {
	layout, _ := layouts.NewPageLayout("L", "mm", 15, 15)
	fpdf := pdf.NewGoFpdf(layout)

	assert.Equal(s.T(), layout, fpdf.Layout())
}

func newPdf() *pdf.GoFpdf {
	layout, _ := layouts.NewPageLayout("L", "mm", 15, 15)
	easypdf := pdf.NewGoFpdf(layout)

	return easypdf
}
