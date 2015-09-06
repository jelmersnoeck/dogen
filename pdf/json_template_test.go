package pdf_test

import (
	"encoding/json"
	"testing"

	"github.com/jelmersnoeck/noscito/pdf"
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
	template, ok := pdf.NewJsonTemplate(inputBytes())

	assert.NotNil(s.T(), template, "Template can't be nil")
	assert.True(s.T(), ok, "Template not created properly")
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

func (s *JsonTemplateSuite) TestBlocks() {
	template, _ := pdf.NewJsonTemplate(inputBytes())

	blocks := template.Blocks()
	assert.Nil(s.T(), blocks)
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
