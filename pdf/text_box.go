package pdf

type TextBox struct {
	Text string  `mapstructure:"text"`
	X    float64 `mapstructure:"x"`
	Y    float64 `mapstructure:"y"`
}

// Parse puts the text on a specific position on the page.
func (b *TextBox) Parse(pdf Pdf) {
	startX, startY := pdf.Position()

	pdf.SetPosition(b.X, b.Y)
	pdf.Text(b.Text)
	pdf.SetPosition(startX, startY)
}

// Load sets all the options for the text block like transformation, font,
// color, etc.
func (i *TextBox) Load(t Template, block_data, user_input map[string]interface{}) {
}
