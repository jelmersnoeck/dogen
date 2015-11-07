package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/dogen/renderer/documents/pdf"
	"github.com/jelmersnoeck/dogen/renderer/mocks"
	"github.com/jelmersnoeck/dogen/renderer/template"
	"github.com/stretchr/testify/suite"
)

type PdfSuite struct {
	suite.Suite
}

func TestPdfSuite(t *testing.T) {
	suite.Run(t, new(PdfSuite))
}

func (s *PdfSuite) TestParseBlocks() {
	block_0 := new(mocks.Block)
	block_1 := new(mocks.Block)
	blocks := make([]template.Block, 2)
	blocks[0] = block_0
	blocks[1] = block_1

	mpdf := &mocks.Pdf{}
	doc := &mocks.Document{}

	mpdf.On("Document").Return(doc)

	block_0.On("Parse", doc).Return(true)
	block_1.On("Parse", doc).Return(true)

	pdf.ParseBlocks(mpdf, blocks)

	block_0.AssertExpectations(s.T())
	block_1.AssertExpectations(s.T())
}
