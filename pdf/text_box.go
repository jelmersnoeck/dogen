package pdf

import "github.com/jelmersnoeck/noscito/utils"

// TextBox takes a string of text and puts it on the given position on the page.
type TextBox struct {
	Text       string   `mapstructure:"text"`
	HTML       string   `mapstructure:"html"`
	Size       float64  `mapstructure:"size"`
	Color      string   `mapstructure:"color"`
	Width      float64  `mapstructure:"width"`
	Fill       bool     `mapstructure:"fill"`
	Align      string   `mapstructure:"align"`
	LineHeight float64  `mapstructure:"lineheight"`
	Position   Position `mapstructure:"position"`
}

// Parse puts the text on a specific position on the page.
func (b *TextBox) Parse(doc Document) {
	doc.AddFont("ProximaNova", "", "proximanova-regular-webfont.json")
	doc.AddFont("ProximaNova", "B", "proximanova-bold-webfont.json")

	doc.SetTextColor(utils.HexToRGB(b.Color))
	doc.SetFont("ProximaNova", "", b.Size)

	if b.Text != "" {
		b.setPosition(doc)
		doc.MultiCell(b.Width, doc.PointConvert(b.LineHeight), b.Text, "", b.Align, b.Fill)
	}

	if b.HTML != "" {
		leftMargin, topMargin, rightMargin, _ := doc.GetMargins()
		doc.SetLeftMargin(leftMargin + b.Position.X)
		doc.SetY(b.Position.Y + topMargin)

		if b.Width > 0 {
			pageWidth, _ := doc.GetPageSize()
			rm := pageWidth - rightMargin - (b.Position.X + b.Width)
			doc.SetRightMargin(rm)
		}

		html := doc.HTMLBasicNew()
		html.Write(doc.PointConvert(b.LineHeight), b.HTML)

		doc.SetLeftMargin(leftMargin)
		doc.SetRightMargin(rightMargin)
	}
}

// Load sets all the options for the text block like transformation, font,
// color, etc.
func (b *TextBox) Load(t Template, block_data, user_input map[string]interface{}) {
	if block_data["align"] == nil {
		b.Align = "CM"
	}

	if block_data["lineheight"] == nil {
		b.LineHeight = b.Size + 2
	}
}

func (b *TextBox) setPosition(doc Document) {
	if b.Position.Y >= 0 {
		doc.SetY(b.Position.Y)
	}

	if b.Position.X >= 0 {
		doc.SetX(b.Position.X)
	}
}
