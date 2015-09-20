package pdf

type AddPage struct {
	Margin Margin `mapstructure:"margin"`
}

func (b *AddPage) Parse(doc Document) {
	doc.AddPage()
	doc.SetMargins(b.Margin.Left, b.Margin.Top, b.Margin.Right)
	doc.SetAutoPageBreak(true, b.Margin.Bottom)
}

func (b *AddPage) Load(t Template, block_data, user_input map[string]interface{}) {
}
