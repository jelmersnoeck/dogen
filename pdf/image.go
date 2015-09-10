package pdf

import "github.com/jung-kurt/gofpdf/contrib/httpimg"

type Image struct {
	Url string  `mapstructure:"url"`
	X   float64 `mapstructure:"x"`
	Y   float64 `mapstructure:"y"`
	W   float64 `mapstructure:"width"`
	H   float64 `mapstructure:"height"`
}

// Parse puts a HTTP Image on the PDF.
func (i *Image) Parse(doc Document) {
	httpimg.Register(doc, i.Url, "")
	doc.Image(i.Url, i.X, i.Y, i.W, i.H, false, "", 0, "")
}

// Load will look if there is a url key in the user input and if so, it will
// overwrite the given URL from the template with this URL.
func (i *Image) Load(t Template, block_data, user_input map[string]interface{}) {
	if url, ok := user_input["url"]; ok {
		i.Url = url.(string)
	}
}
