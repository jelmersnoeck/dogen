package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/assert"
)

func TestNewPageSizeDefaults(t *testing.T) {
	s, err := pdf.NewPageSize("", "", 1, 1)

	assert.Equal(t, "L", s.Orientation())
	assert.Equal(t, "mm", s.Unit())
	assert.Nil(t, err)
}

func TestWidthError(t *testing.T) {
	_, err := pdf.NewPageSize("", "", 0, 1)

	assert.NotNil(t, err)
}

func TestHeightError(t *testing.T) {
	_, err := pdf.NewPageSize("", "", 1, 0)

	assert.NotNil(t, err)
}

func TestWidth(t *testing.T) {
	s, _ := pdf.NewPageSize("", "", 15, 20)

	assert.EqualValues(t, 15, s.Width())
}

func TestHeight(t *testing.T) {
	s, _ := pdf.NewPageSize("", "", 15, 20)

	assert.EqualValues(t, 20, s.Height())
}
