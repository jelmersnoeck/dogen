package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
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

func (s *GoFpdfSuite) TestParseBlocks() {
	block_0 := new(mocks.Block)
	block_1 := new(mocks.Block)
	blocks := make([]pdf.Block, 2)
	blocks[0] = block_0
	blocks[1] = block_1

	f := newPdf()
	block_0.On("Parse", f).Return(true)
	block_1.On("Parse", f).Return(true)

	f.ParseBlocks(blocks)

	block_0.AssertExpectations(s.T())
	block_1.AssertExpectations(s.T())
}

func newPdf() *pdf.GoFpdf {
	layout, _ := pdf.NewPageLayout("L", "mm", 15, 15)
	easypdf := pdf.NewGoFpdf(layout)

	return easypdf
}
