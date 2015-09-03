package pdf_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
	"github.com/stretchr/testify/assert"
)

func TestNewPdf(t *testing.T) {
	easypdf := newPdf()

	var ep *pdf.EasyPdf
	assert.IsType(t, ep, easypdf)
}

func TestBytesWithoutError(t *testing.T) {
	easypdf := newPdf()

	drawer := new(mocks.PdfDrawer)
	easypdf.Drawer = drawer

	buffer := bytes.NewBufferString("")
	drawer.On("Output", buffer).Return(nil)

	assert.NotNil(t, easypdf.Bytes(buffer))
}

func TestBytesWithError(t *testing.T) {
	easypdf := newPdf()

	drawer := new(mocks.PdfDrawer)
	easypdf.Drawer = drawer

	buffer := bytes.NewBufferString("")
	err := errors.New("Error")
	drawer.On("Output", buffer).Return(err)
	drawer.On("SetError", err).Return(err)

	assert.Nil(t, easypdf.Bytes(buffer))
}

func newPdf() *pdf.EasyPdf {
	size, _ := pdf.NewPageSize("L", "mm", 15, 15)
	easypdf := pdf.NewPdf(size)

	return easypdf
}
