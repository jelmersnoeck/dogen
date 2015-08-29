package easypdf

import ()

type UserInput struct {
	InputId string `mapstructure:"input_id"`
	Block   Block  `mapstructure:"block"`
}

func (ui *UserInput) Parse(pdf *EasyPDF, input map[string]interface{}) {
	data := input[ui.InputId].(map[string]interface{})
	ui.Block.Parse(pdf, data)
}

func (ui *UserInput) Load(data interface{}) {
	mapped := data.(map[string]interface{})
	ui.InputId = mapped["input_id"].(string)
	ui.Block = LoadBlock(mapped["block"])
}
