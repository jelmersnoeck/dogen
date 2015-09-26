package pdf

import "github.com/jelmersnoeck/noscito/utils"

type FontType struct {
	Size       float64 `mapstructure:"size"`
	Color      string  `mapstructure:"color"`
	LineHeight float64 `mapstructure:"lineheight"`
	Weight     string  `mapstructure:"weight"`
	Type       string  `mapstructure:"type"`
}

func (f FontType) Register(doc Document) {
	if f.Color != "" {
		doc.SetTextColor(utils.HexToRGB(f.Color))
	}

	if f.Type != "" && f.Size != 0 {
		doc.SetFont(f.Type, f.Weight, f.Size)
	}
}
