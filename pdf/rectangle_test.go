package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/suite"
)

type RectangleBlockSuite struct {
	suite.Suite
}

func TestRectangleBlockSuite(t *testing.T) {
	suite.Run(t, new(RectangleBlockSuite))
}

func (s *RectangleBlockSuite) TestParse() {
	doc := &mocks.Document{}
	rect := &pdf.Rectangle{"ff00ff", 50, 60, pdf.Position{1, 5}, 0}

	doc.On("SetFillColor", 255, 00, 255).Return()
	doc.On("Rect", 1.0, 5.0, 50.0, 60.0, "F").Return()

	rect.Parse(doc)

	doc.AssertExpectations(s.T())
}
