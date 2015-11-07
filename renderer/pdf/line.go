package pdf

import "github.com/jelmersnoeck/dogen/renderer/utils"

type Line struct {
	Color    string   `mapstructure:"color"`
	Position Position `mapstructure:"position"`
	W        float64  `mapstructure:"width"`
	H        float64  `mapstructure:"height"`
}

func (b *Line) Parse(doc Document) {
	doc.SetDrawColor(utils.HexToRGB(b.Color))
	doc.Line(b.Position.X, b.Position.Y, b.Position.X+b.W, b.Position.Y+b.H)
}

func (b *Line) Load(t Template, block_data, user_input map[string]interface{}) {
}
