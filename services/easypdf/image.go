package easypdf

import (
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

type Image struct {
	Url string  `mapstructure:"url"`
	X   float64 `mapstructure:"x"`
	Y   float64 `mapstructure:"y"`
	W   float64 `mapstructure:"w"`
	H   float64 `mapstructure:"h"`
}

func (i *Image) Parse(pdf *EasyPDF, input map[string]interface{}) {
	i.Url = i.urlFromInput(input)

	pdf.Wg.Add(1)
	go putImage(pdf, i)
}

func (i *Image) Load(data interface{}) {
}

func (i *Image) urlFromInput(input map[string]interface{}) string {
	return input["url"].(string)
}

func putImage(pdf *EasyPDF, i *Image) {
	defer pdf.Wg.Done()

	httpimg.Register(pdf.pdf, i.Url, "")
	pdf.pdf.Image(i.Url, i.X, i.Y, i.W, i.H, false, "", 0, "")
}
