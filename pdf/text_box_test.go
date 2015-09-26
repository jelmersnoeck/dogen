package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/suite"
)

type TextBoxBlockSuite struct {
	suite.Suite
}

func TestTextBoxBlockSuite(t *testing.T) {
	suite.Run(t, new(TextBoxBlockSuite))
}

func (s *TextBoxBlockSuite) TestFontSettings() {
	doc, txtBx := setupDocument("", "")

	txtBx.Parse(doc)

	doc.AssertExpectations(s.T())
}

func (s *TextBoxBlockSuite) TestTextParsing() {
	doc, txtBx := setupDocument("Test document", "")

	doc.On("SetY", 10.0).Return(true)
	doc.On("SetX", 5.0).Return(true)
	doc.On("PointConvert", 12.0).Return(20.0)
	doc.On("MultiCell", 5.5, 20.0, "Test document", "", "C", false).Return(true)

	txtBx.Parse(doc)

	doc.AssertExpectations(s.T())
}

func (s *TextBoxBlockSuite) TestHTMLParsing() {
	doc, txtBx := setupDocument("", "HTML Document")
	pdf := gofpdf.New("P", "mm", "A4", "")

	doc.On("GetMargins").Return(1.0, 2.0, 3.0, 4.0)
	doc.On("GetPageSize").Return(20.0, 30.0)
	doc.On("PointConvert", 12.0).Return(20.0)

	doc.On("SetLeftMargin", 6.0).Return(true)
	doc.On("SetLeftMargin", 1.0).Return(true)
	doc.On("SetRightMargin", 6.5).Return(true)
	doc.On("SetRightMargin", 3.0).Return(true)
	doc.On("SetY", 12.0).Return(true)
	doc.On("HTMLBasicNew").Return(pdf.HTMLBasicNew())

	txtBx.Parse(doc)

	doc.AssertExpectations(s.T())
}

func setupDocument(txtStr, htmlStr string) (*mocks.Document, *pdf.TextBox) {
	doc := &mocks.Document{}
	txtBx := &pdf.TextBox{
		txtStr,
		htmlStr,
		pdf.FontType{12, "ffffff", 12, "", "Helvetica"},
		5.5,
		false,
		"C",
		pdf.Position{5, 10},
	}

	doc.On("SetTextColor", 255, 255, 255).Return()
	doc.On("SetFont", "Helvetica", "", 12.0).Return()

	return doc, txtBx
}
