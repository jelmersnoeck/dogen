package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/assert"
)

func TestNewGoFpdf(t *testing.T) {
	easypdf := newPdf()

	var ep *pdf.GoFpdf
	assert.IsType(t, ep, easypdf)
}

func TestParseBlocks(t *testing.T) {
	block_0 := new(mocks.Block)
	block_1 := new(mocks.Block)
	blocks := make([]pdf.Block, 2)
	blocks[0] = block_0
	blocks[1] = block_1

	f := newPdf()
	block_0.On("Parse", f).Return(true)
	block_1.On("Parse", f).Return(true)

	f.ParseBlocks(blocks)

	block_0.AssertExpectations(t)
	block_1.AssertExpectations(t)
}

func newPdf() *pdf.GoFpdf {
	layout, _ := pdf.NewPageLayout("L", "mm", 15, 15)
	easypdf := pdf.NewGoFpdf(layout)

	return easypdf
}
