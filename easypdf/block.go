package easypdf

import (
	"github.com/mitchellh/mapstructure"
)

type Block interface {
	Parse(pdf *EasyPDF, input map[string]interface{})
}

func NewBlock(item interface{}) (block Block) {
	block = getBlock(item)
	mapstructure.Decode(item, block)

	return
}

func getBlock(block interface{}) (b Block) {
	data := block.(map[string]interface{})
	switch getType(data) {
	case "user_input":
		b := new(UserInput)
		b.Load(data)
		return b
	case "image":
		return new(Image)
	}
	return nil
}

func getType(data map[string]interface{}) string {
	return data["type"].(string)
}

func MappedData(input interface{}) map[string]interface{} {
	return input.(map[string]interface{})
}
