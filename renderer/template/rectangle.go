package template

import (
	"fmt"

	"github.com/jelmersnoeck/dogen/renderer/documents"
	"github.com/jelmersnoeck/dogen/renderer/utils"
)

// TextBox takes a string of text and puts it on the given position on the page.
type Rectangle struct {
	Color    string   `mapstructure:"color"`
	Width    float64  `mapstructure:"width"`
	Height   float64  `mapstructure:"height"`
	Position Position `mapstructure:"position"`
	Rotation float64
}

// Parse puts the text on a specific position on the page.
func (b *Rectangle) Parse(doc documents.Document) {
	if b.Rotation != 0 {
		doc.TransformBegin()
		fmt.Println(b.Position.X + b.Width/2)
		fmt.Println(b.Position.Y + b.Height/2)
		doc.TransformRotate(b.Rotation, b.Position.X+b.Width/2, b.Position.Y+b.Height/2)
	}

	doc.SetFillColor(utils.HexToRGB(b.Color))
	doc.Rect(b.Position.X, b.Position.Y, b.Width, b.Height, "F")

	if b.Rotation != 0 {
		doc.TransformEnd()
	}
}

// Load sets all the options for the text block like transformation, font,
// color, etc.
func (i *Rectangle) Load(t Template, block_data, user_input map[string]interface{}) {
}
