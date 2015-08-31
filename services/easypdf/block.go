package easypdf

import (
	"github.com/mitchellh/mapstructure"
)

type Block interface {
	Parse(pdf *EasyPDF, input map[string]interface{})
	Load(data interface{})
}

func NewBlock(item interface{}) (block Block) {
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
