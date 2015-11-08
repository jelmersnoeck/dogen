package templates

import (
	"strings"

	"github.com/jelmersnoeck/dogen/renderer/documents"
)

// TextBox takes a string of text and puts it on the given position on the page.
type TextBox struct {
	Text     string   `mapstructure:"text"`
	HTML     string   `mapstructure:"html"`
	Font     FontType `mapstructure:"font"`
	Width    float64  `mapstructure:"width"`
	Fill     bool     `mapstructure:"fill"`
	Align    string   `mapstructure:"align"`
	Position Position `mapstructure:"position"`
	Rotation float64
	Height   float64
}

// Parse puts the text on a specific position on the page.
func (b *TextBox) Parse(doc documents.Document) {
	b.Font.Register(doc)

	if b.Rotation != 0 {
		doc.TransformBegin()

		rotationX, rotationY := b.getRotation(doc)
		doc.TransformRotate(b.Rotation, rotationX, rotationY)
	}

	if b.Text != "" {
		b.setPosition(doc)
		doc.MultiCell(b.Width, doc.PointConvert(b.Font.LineHeight), b.Text, "", b.Align, b.Fill)
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
		html.Write(doc.PointConvert(b.Font.LineHeight), b.HTML)

		doc.SetLeftMargin(leftMargin)
		doc.SetRightMargin(rightMargin)
	}

	if b.Rotation != 0 {
		doc.TransformEnd()
	}
}

// Load sets all the options for the text block like transformation, font,
// color, etc.
func (b *TextBox) Load(t Template, block_data, user_input map[string]interface{}) {
	if block_data["align"] == nil {
		b.Align = "CM"
	}

	if block_data["lineheight"] == nil {
		b.Font.LineHeight = b.Font.Size + 2
	}

	if block_data["replace"] != nil {
		for _, value := range block_data["replace"].([]interface{}) {
			key := string(value.(string))
			replacer := strings.NewReplacer("%{"+key+"}%", user_input[key].(string))
			b.HTML = replacer.Replace(b.HTML)
		}
	}

	if text, ok := user_input["text"]; ok {
		b.Text = text.(string)
	}
}

func (b *TextBox) setPosition(doc documents.Document) {
	if b.Position.Y >= 0 {
		doc.SetY(b.Position.Y)
	}

	if b.Position.X >= 0 {
		doc.SetX(b.Position.X)
	}
}

func (b *TextBox) getRotation(doc documents.Document) (rotationX float64, rotationY float64) {
	if b.Width == 0 {
		pageWidth, _ := doc.GetPageSize()
		leftMargin, _, rightMargin, _ := doc.GetMargins()

		rotationX = (pageWidth - leftMargin - rightMargin - b.Position.X) / 2
	} else {
		rotationX = b.Position.X + b.Width/2
	}

	if b.Height == 0 {
		rotationY = b.Position.Y + doc.PointConvert(b.Font.LineHeight)/2
	} else {
		rotationY = b.Position.Y + b.Height/2
	}

	return
}
