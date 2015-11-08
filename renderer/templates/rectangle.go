package templates

import (
	"fmt"

	"github.com/jelmersnoeck/dogen/renderer/documents"
	"github.com/jelmersnoeck/dogen/renderer/utils"
)

// Rectangle represents a rectangular box with borders that gets drawn on the
// document.
type Rectangle struct {
	Color    string   `mapstructure:"color"`
	Width    float64  `mapstructure:"width"`
	Height   float64  `mapstructure:"height"`
	Position Position `mapstructure:"position"`
	Rotation float64
}

// Parse will draw the rectangle on the document with the given attributes.
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

// Load is currently not used for a Rectangle.
func (i *Rectangle) Load(t Template, block_data, user_input map[string]interface{}) {
}
