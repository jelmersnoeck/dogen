package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/pdf"
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

func newPdf() *pdf.GoFpdf {
	layout, _ := pdf.NewPageLayout("L", "mm", 15, 15)
	easypdf := pdf.NewGoFpdf(layout)

	return easypdf
}
