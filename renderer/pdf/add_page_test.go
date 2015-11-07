package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/renderer/mocks"
	"github.com/jelmersnoeck/noscito/renderer/pdf"
	"github.com/stretchr/testify/suite"
)

type AddPageBlockSuite struct {
	suite.Suite
}

func TestAddPageBlockSuite(t *testing.T) {
	suite.Run(t, new(AddPageBlockSuite))
}

func (s *AddPageBlockSuite) TestParse() {
	doc := &mocks.Document{}
	page := &pdf.AddPage{pdf.Margin{5, 10, 15, 20}}

	doc.On("AddPage").Return()
	doc.On("SetMargins", 5.0, 10.0, 15.0).Return()
	doc.On("SetAutoPageBreak", true, 20.0).Return()

	page.Parse(doc)

	doc.AssertExpectations(s.T())
}
