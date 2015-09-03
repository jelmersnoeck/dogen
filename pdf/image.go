package pdf

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

func (i *Image) Parse(pdf Pdf) {
	httpimg.Register(pdf.Drawer(), i.Url, "")
	pdf.Drawer().Image(i.Url, i.X, i.Y, i.W, i.H, false, "", 0, "")
}
