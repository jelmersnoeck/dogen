package pdf_test

import (
	"encoding/json"
	"testing"

	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/assert"
)

func TestNewJsonTemplate(t *testing.T) {
	template, ok := pdf.NewJsonTemplate(inputBytes())

	assert.NotNil(t, template, "Template can't be nil")
	assert.True(t, ok, "Template not created properly")
}

func TestLayout(t *testing.T) {
	template, _ := pdf.NewJsonTemplate(inputBytes())

	layout := template.Layout()
	assert.NotNil(t, layout, "Layout can't be nil")
	assert.EqualValues(t, "L", layout.Orientation(), "Orientation is not correct")
	assert.EqualValues(t, "mm", layout.Unit(), "Unit is not correct")
	assert.EqualValues(t, 50, layout.Width(), "Width is not correct")
	assert.EqualValues(t, 75, layout.Height(), "Height is not correct")
}

func TestLoadBlocks(t *testing.T) {
	template, _ := pdf.NewJsonTemplate(inputBytes())

	template.LoadBlocks(userInput())
}

func TestBlocks(t *testing.T) {
	template, _ := pdf.NewJsonTemplate(inputBytes())

	blocks := template.Blocks()
	assert.Nil(t, blocks)
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
				"block_data": {
					"x": 10,
					"y": 10,
					"w": 50,
					"h": 50,
					"url": "https://upload.wikimedia.org/wikipedia/commons/thumb/a/aa/Logo_Google_2013_Official.svg/1280px-Logo_Google_2013_Official.svg.png"
				}
			}
		]
	}`)
}

func userInput() (data map[string]interface{}) {
	byt := []byte(`{
		"set_1_image_1": { "url": "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/UPC-A-036000291452.png/220px-UPC-A-036000291452.png", "visible": false },
		"set_1_image_2": { "url": "http://petapixel.com/assets/uploads/2013/11/bloomf1.jpeg" },
		"loop": {
			"items": [
			]
		}
	}`)

	json.Unmarshal(byt, &data)
	return data
}
