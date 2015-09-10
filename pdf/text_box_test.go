package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/stretchr/testify/suite"
)

type TextBoxBlockSuite struct {
	suite.Suite
}

func TestTextBoxBlockSuite(t *testing.T) {
	suite.Run(t, new(TextBoxBlockSuite))
}

func (s *TextBloxBlockSuite) TestPositioning() {
	mpdf := &mocks.Pdf{}
}
