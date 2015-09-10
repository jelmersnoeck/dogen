package pdf

// TextBox takes a string of text and puts it on the given position on the page.
type TextBox struct {
	Text string  `mapstructure:"text"`
	X    float64 `mapstructure:"x"`
	Y    float64 `mapstructure:"y"`
}

// Parse puts the text on a specific position on the page.
func (b *TextBox) Parse(doc Document) {
	startX, startY := doc.GetXY()

	doc.SetXY(b.X, b.Y)
	doc.MultiCell(0, 5, b.Text, "", "", false)
	doc.SetXY(startX, startY)
}

// Load sets all the options for the text block like transformation, font,
// color, etc.
func (i *TextBox) Load(t Template, block_data, user_input map[string]interface{}) {
}
