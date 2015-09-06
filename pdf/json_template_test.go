package pdf_test

import (
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

func inputBytes() []byte {
	return []byte(`{
		"layout": {
			"orientation": "L",
			"unit": "mm",
			"width": 50,
			"height": 75
		}
	}`)
}
