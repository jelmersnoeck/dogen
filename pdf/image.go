package pdf

type Image struct {
	Url string  `mapstructure:"url"`
	X   float64 `mapstructure:"x"`
	Y   float64 `mapstructure:"y"`
	W   float64 `mapstructure:"w"`
	H   float64 `mapstructure:"h"`
}

// Parse puts a HTTP Image on the PDF.
func (i *Image) Parse(pdf Pdf) {
	pdf.HttpImage(i.Url, i.X, i.Y, i.W, i.H, false, "", 0, "")
}
