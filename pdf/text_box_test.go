package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/suite"
)

type TextBoxBlockSuite struct {
	suite.Suite
}

func TestTextBoxBlockSuite(t *testing.T) {
	suite.Run(t, new(TextBoxBlockSuite))
}

func (s *TextBoxBlockSuite) TestPositioning() {
	doc := &mocks.Document{}
	txtBx := &pdf.TextBox{"Text", 5.5, 5.5}

	doc.On("GetXY").Return(10.5, 10.5)
	doc.On("SetXY", 5.5, 5.5).Return(true)
	doc.On("SetXY", 10.5, 10.5).Return(true)
	doc.On("MultiCell", 0.0, 5.0, txtBx.Text, "", "", false).Return(true)

	txtBx.Parse(doc)

	doc.AssertExpectations(s.T())
}
