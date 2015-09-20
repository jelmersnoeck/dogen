package pdf

type FontType struct {
	Size       float64 `mapstructure:"size"`
	Color      string  `mapstructure:"color"`
	LineHeight float64 `mapstructure:"lineheight"`
	Weight     string  `mapstructure"weight"`
}
