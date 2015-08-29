package block

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

type Image struct {
	Url string  `mapstructure:"url"`
	X   float64 `mapstructure:"x"`
	Y   float64 `mapstructure:"y"`
	W   float64 `mapstructure:"w"`
	H   float64 `mapstructure:"h"`
}

func (i *Image) Parse(pdf *gofpdf.Fpdf, input map[string]interface{}) {
	i.Url = i.urlFromInput(input)

	httpimg.Register(pdf, i.Url, "")
	pdf.Image(i.Url, i.X, i.Y, i.W, i.H, false, "", 0, "")
}

func (i *Image) Load(data interface{}) {
}

func (i *Image) urlFromInput(input map[string]interface{}) string {
	return input["url"].(string)
}
