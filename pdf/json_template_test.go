package pdf_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/jelmersnoeck/noscito/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JsonTemplateSuite struct {
	suite.Suite
}

func TestJsonTemplateSuite(t *testing.T) {
	suite.Run(t, new(JsonTemplateSuite))
}

func (s *JsonTemplateSuite) TestNewJsonTemplate() {
	template, err := pdf.NewJsonTemplate(inputBytes())

	assert.NotNil(s.T(), template, "Template can't be nil")
	assert.Nil(s.T(), err, "Template not created properly")
}

func (s *JsonTemplateSuite) TestLayout() {
	template, _ := pdf.NewJsonTemplate(inputBytes())

	layout := template.Layout()
	assert.NotNil(s.T(), layout, "Layout can't be nil")
	assert.EqualValues(s.T(), "L", layout.Orientation(), "Orientation is not correct")
	assert.EqualValues(s.T(), "mm", layout.Unit(), "Unit is not correct")
	assert.EqualValues(s.T(), 50, layout.Width(), "Width is not correct")
	assert.EqualValues(s.T(), 75, layout.Height(), "Height is not correct")
}

func (s *JsonTemplateSuite) TestLoadBlock() {
	template, _ := pdf.NewJsonTemplate(inputBytes())

	template.LoadBlocks(userInput())

	blocks := template.Blocks()
	assert.NotNil(s.T(), blocks[0])

	var img *pdf.Image
	assert.IsType(s.T(), img, blocks[0])
}

func (s *JsonTemplateSuite) TestErrors() {
	template := &pdf.JsonTemplate{}

	assert.Equal(s.T(), 0, len(template.Errors()))

	err := errors.New("New error")
	template.AddError(err)

	assert.Equal(s.T(), template.Errors()[0], err)
}

func (s *JsonTemplateSuite) TestBlocks() {
	template, _ := pdf.NewJsonTemplate(inputBytes())

	blocks := template.Blocks()
	assert.Nil(s.T(), blocks)
}

func (s *JsonTemplateSuite) TestLoadImage() {
	template_data, _ := utils.LoadTemplate("pb-collection")
	template, _ := pdf.NewJsonTemplate(template_data)
	block_data := map[string]interface{}{"type": "image", "url": "https://upload.wikimedia.org/wikipedia/commons/thumb/1/1e/Wikipedia-logo-v2-es.svg/2000px-Wikipedia-logo-v2-es.svg.png"}
	input_data := map[string]interface{}{}

	block := template.LoadBlock(block_data, input_data)

	var img *pdf.Image
	assert.IsType(s.T(), img, block)
}

func (s *JsonTemplateSuite) TestLoadRectangle() {
	template_data, _ := utils.LoadTemplate("pb-collection")
	template, _ := pdf.NewJsonTemplate(template_data)
	block_data := map[string]interface{}{"type": "rectangle"}
	input_data := map[string]interface{}{}

	block := template.LoadBlock(block_data, input_data)

	var rct *pdf.Rectangle
	assert.IsType(s.T(), rct, block)
}

func (s *JsonTemplateSuite) TestLoadLine() {
	template_data, _ := utils.LoadTemplate("pb-collection")
	template, _ := pdf.NewJsonTemplate(template_data)
	block_data := map[string]interface{}{"type": "line"}
	input_data := map[string]interface{}{}

	block := template.LoadBlock(block_data, input_data)

	var rct *pdf.Line
	assert.IsType(s.T(), rct, block)
}

func (s *JsonTemplateSuite) TestLoadAddPage() {
	template_data, _ := utils.LoadTemplate("pb-collection")
	template, _ := pdf.NewJsonTemplate(template_data)
	block_data := map[string]interface{}{"type": "add_page"}
	input_data := map[string]interface{}{}

	block := template.LoadBlock(block_data, input_data)

	var pg *pdf.AddPage
	assert.IsType(s.T(), pg, block)
}

func (s *JsonTemplateSuite) TestLoadUserInput() {
	template := &pdf.JsonTemplate{}
	// block made optional so we don't need input data
	block_data := map[string]interface{}{"type": "user_input", "optional": true}
	input_data := map[string]interface{}{}

	block := template.LoadBlock(block_data, input_data)

	var ui *pdf.UserInput
	assert.IsType(s.T(), ui, block)
}

func (s *JsonTemplateSuite) TestLoadTextBox() {
	template := &pdf.JsonTemplate{}
	block_data := map[string]interface{}{"type": "text_box", "text": "Jelmer"}
	input_data := map[string]interface{}{}

	block := template.LoadBlock(block_data, input_data)

	var txt *pdf.TextBox
	assert.IsType(s.T(), txt, block)
}

func inputBytes() []byte {
	return []byte(`{
		"layout": {
			"orientation": "L",
			"unit": "mm",
			"width": 50,
			"height": 75
		},
		"blocks": [
			{
				"type": "image",
				"x": 10,
				"y": 10,
				"w": 50,
				"h": 50,
				"url": "https://upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Logo_Google_2013_Official.svg/1280px-Logo_Google_2013_Official.svg.png"
			}
		]
	}`)
}

func userInput() (data map[string]interface{}) {
	byt := []byte(`{}`)
	json.Unmarshal(byt, &data)
	return data
}
