package pdf_test

import (
	"bytes"
	"testing"

	"github.com/jelmersnoeck/noscito/renderer/mocks"
	"github.com/jelmersnoeck/noscito/renderer/pdf"
	"github.com/jelmersnoeck/noscito/renderer/utils"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ImageBlockSuite struct {
	suite.Suite
}

func TestImageBlockSuite(t *testing.T) {
	suite.Run(t, new(ImageBlockSuite))
}

func (s *ImageBlockSuite) TestParse() {
	reader := bytes.NewReader([]byte{})
	img := &pdf.Image{"url", 20, 20, 50, 50, "image/jpeg", reader}
	doc := &mocks.Document{}

	d := float64(20)
	p := float64(50)

	doc.On("ImageTypeFromMime", "image/jpeg").Return("jpeg")
	doc.On("RegisterImageReader", "url", "jpeg", reader).Return(&gofpdf.ImageInfoType{})
	doc.On("Image", "url", d, d, p, p, false, "jpeg", 0, "").Return(true)

	img.Parse(doc)

	doc.AssertExpectations(s.T())
}

func (s *ImageBlockSuite) TestLoadNoOverwrite() {
	reader := bytes.NewReader([]byte{})
	img := &pdf.Image{"url", 20, 20, 50, 50, "image/jpeg", reader}
	template_data, _ := utils.LoadTemplate("test-pdf")
	template, _ := pdf.NewJsonTemplate(template_data)

	block_data := map[string]interface{}{}
	user_input := map[string]interface{}{}

	img.Load(template, block_data, user_input)

	assert.EqualValues(s.T(), "url", img.Url)
}

func (s *ImageBlockSuite) TestLoadOverwrite() {
	reader := bytes.NewReader([]byte{})
	img := &pdf.Image{"url", 20, 20, 50, 50, "image/jpeg", reader}
	template_data, _ := utils.LoadTemplate("test-pdf")
	template, _ := pdf.NewJsonTemplate(template_data)

	block_data := map[string]interface{}{}
	user_input := map[string]interface{}{"url": "new-url"}

	img.Load(template, block_data, user_input)

	assert.EqualValues(s.T(), "new-url", img.Url)
}
