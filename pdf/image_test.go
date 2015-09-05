package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/mocks"
	"github.com/jelmersnoeck/noscito/pdf"
)

func TestParse(t *testing.T) {
	img := &pdf.Image{"url", 20, 20, 50, 50}
	mpdf := &mocks.MPdf{}

	d := float64(20)
	p := float64(50)
	mpdf.On("HttpImage", "url", d, d, p, p, false, "", 0, "").Return(true)

	img.Parse(mpdf)

	mpdf.AssertExpectations(t)
}
