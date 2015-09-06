package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
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
	img := &pdf.Image{"url", 20, 20, 50, 50}
	mpdf := &mocks.MPdf{}

	d := float64(20)
	p := float64(50)
	mpdf.On("HttpImage", "url", d, d, p, p, false, "", 0, "").Return(true)

	img.Parse(mpdf)

	mpdf.AssertExpectations(s.T())
}

func (s *ImageBlockSuite) TestLoadNoOverwrite() {
	img := &pdf.Image{"url", 20, 20, 50, 50}
	template := &mocks.Template{}

	block_data := map[string]interface{}{}
	user_input := map[string]interface{}{}

	img.Load(template, block_data, user_input)

	assert.EqualValues(s.T(), "url", img.Url)
}

func (s *ImageBlockSuite) TestLoadOverwrite() {
	img := &pdf.Image{"url", 20, 20, 50, 50}
	template := &mocks.Template{}

	block_data := map[string]interface{}{}
	user_input := map[string]interface{}{"url": "new-url"}

	img.Load(template, block_data, user_input)

	assert.EqualValues(s.T(), "new-url", img.Url)
}
