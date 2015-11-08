package templates

import "github.com/jelmersnoeck/dogen/renderer/documents"

// AddPage is a BlockType that allows to add a new page on the document.
type AddPage struct {
	Margin Margin `mapstructure:"margin"`
}

// Parse will add a new page to the document and set the appropiate page
// settings.
func (b *AddPage) Parse(doc documents.Document) {
	doc.AddPage()
	doc.SetMargins(b.Margin.Left, b.Margin.Top, b.Margin.Right)
	doc.SetAutoPageBreak(true, b.Margin.Bottom)
}

// Load currently does nothing for AddPage.
func (b *AddPage) Load(t Template, block_data, user_input map[string]interface{}) {
}
