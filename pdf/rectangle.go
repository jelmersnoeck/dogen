package pdf

import "github.com/jelmersnoeck/noscito/utils"

// TextBox takes a string of text and puts it on the given position on the page.
type Rectangle struct {
	Color    string   `mapstructure:"color"`
	Width    float64  `mapstructure:"width"`
	Height   float64  `mapstructure:"height"`
	Position Position `mapstructure:"position"`
}

// Parse puts the text on a specific position on the page.
func (b *Rectangle) Parse(doc Document) {
	doc.SetFillColor(utils.HexToRGB(b.Color))
	doc.Rect(b.Position.X, b.Position.Y, b.Width, b.Height, "F")
}

// Load sets all the options for the text block like transformation, font,
// color, etc.
func (i *Rectangle) Load(t Template, block_data, user_input map[string]interface{}) {
}
