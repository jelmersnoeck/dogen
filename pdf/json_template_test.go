package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/assert"
)

func TestNewJsonTemplate(t *testing.T) {
	bytes := []byte(`{
		"layout": {
			"orientation": "L",
			"unit": "mm",
			"width": 50,
			"height": 50
		}
	}`)
	template, ok := pdf.NewJsonTemplate(bytes)
	assert.True(t, ok)
	assert.NotNil(t, template.Layout())
}
