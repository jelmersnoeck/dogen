package easypdf

import ()

type UserInput struct {
	InputId string `mapstructure:"input_id"`
	Block   Block  `mapstructure:"block"`
}

func (ui *UserInput) Parse(pdf *EasyPDF, input map[string]interface{}) {
	ui.Block.Parse(pdf, ui.inputData(input))
}

func (ui *UserInput) Load(input interface{}) {
	mapped := ui.mappedData(input)
	ui.InputId = mapped["input_id"].(string)
	ui.Block = NewBlock(mapped["block"])
}

func (ui *UserInput) inputData(input map[string]interface{}) map[string]interface{} {
	return ui.mappedData(input[ui.InputId])
}

func (ui *UserInput) mappedData(input interface{}) map[string]interface{} {
	return input.(map[string]interface{})
}
