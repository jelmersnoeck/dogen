package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/suite"
)

type LineBlockSuite struct {
	suite.Suite
}

func TestLineBlockSuite(t *testing.T) {
	suite.Run(t, new(LineBlockSuite))
}

func (s *LineBlockSuite) TestParse() {
	doc := &mocks.Document{}
	line := &pdf.Line{"ff00ff", pdf.Position{5, 10}, 30, 40}

	doc.On("SetDrawColor", 255, 00, 255).Return(true)
	doc.On("Line", 5.0, 10.0, 35.0, 50.0).Return(true)

	line.Parse(doc)

	doc.AssertExpectations(s.T())
}
