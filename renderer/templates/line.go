package templates

import (
	"github.com/jelmersnoeck/dogen/renderer/documents"
	"github.com/jelmersnoeck/dogen/renderer/utils"
)

// Line represents a BlockType to draw a line on the document in a specific
// color.
type Line struct {
	Color    string   `mapstructure:"color"`
	Position Position `mapstructure:"position"`
	W        float64  `mapstructure:"width"`
	H        float64  `mapstructure:"height"`
}

// Parse will draw the line on the document for the given position and color.
func (b *Line) Parse(doc documents.Document) {
	doc.SetDrawColor(utils.HexToRGB(b.Color))
	doc.Line(b.Position.X, b.Position.Y, b.Position.X+b.W, b.Position.Y+b.H)
}

// Load currently does nothing for a Line.
func (b *Line) Load(t Template, block_data, user_input map[string]interface{}) {
}
