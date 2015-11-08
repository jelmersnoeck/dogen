package templates

import (
	"github.com/jelmersnoeck/dogen/renderer/documents"
	"github.com/jelmersnoeck/dogen/renderer/utils"
)

// FontType represents the font we will use in parsing a Block. It holds font
// attributes like Size, Color, LineHeight, Weight and Type.
type FontType struct {
	Size       float64 `mapstructure:"size"`
	Color      string  `mapstructure:"color"`
	LineHeight float64 `mapstructure:"lineheight"`
	Weight     string  `mapstructure:"weight"`
	Type       string  `mapstructure:"type"`
}

// Register will register a font to the document for later usage.
func (f FontType) Register(doc documents.Document) {
	if f.Color != "" {
		doc.SetTextColor(utils.HexToRGB(f.Color))
	}

	if f.Type != "" && f.Size != 0 {
		doc.SetFont(f.Type, f.Weight, f.Size)
	}
}
