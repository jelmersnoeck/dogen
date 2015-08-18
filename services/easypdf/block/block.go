package block

import (
	"encoding/json"
	"github.com/jung-kurt/gofpdf"
)

type BlockItem interface {
	Parse(pdf *gofpdf.Fpdf)
}

type Block struct {
	Type string
	Data json.RawMessage
	Item BlockItem
}

func (b *Block) Unmarshal() {
	switch b.Type {
	case "image":
		i := new(Image)
		json.Unmarshal(b.Data, &i)
		b.Item = i
		b.Data = nil
	}
}
