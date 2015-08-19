package block

import (
	"encoding/json"
	"github.com/jung-kurt/gofpdf"
)

type BlockItem interface {
	Parse(pdf *gofpdf.Fpdf, data map[string]interface{})
}

type Block struct {
	Type string
	Data json.RawMessage
	Item BlockItem
}

func (b *Block) MatchData() {
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
