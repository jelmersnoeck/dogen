package block

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/mitchellh/mapstructure"
)

type Block interface {
	Parse(pdf *gofpdf.Fpdf, input map[string]interface{})
	Load(data interface{})
}

func LoadBlocks(blocks interface{}, stored_blocks []Block) (b []Block) {
	b = make([]Block, len(stored_blocks))
	items := blocks.([]interface{})
	for _, item := range items {
		block := LoadBlock(item)
		b = append(b, block)
	}
	return b
}

func LoadBlock(item interface{}) (block Block) {
	block = getBlock(item)
	mapstructure.Decode(item, block)
	block.Load(item)
	return block
}

func getBlock(block interface{}) (b Block) {
	data := block.(map[string]interface{})
	switch data["type"].(string) {
	case "user_input":
		return new(UserInput)
	case "image":
		return new(Image)
	}
	return nil
}
