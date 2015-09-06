package pdf

type UserInput struct {
	InputId string `mapstructure:"input_id"`
	Block   Block
}

// Parse will recursively parse the blocks that are under this UserInput.
func (b *UserInput) Parse(pdf Pdf) {
	b.Block.Parse(pdf)
}

// Load will recursively load blocks that fall under this block. It will select
// the correct input from the user_input based on the input_id that is specified
// in the block_data.
func (b *UserInput) Load(t Template, block_data, user_input map[string]interface{}) {
	t.LoadBlock(block_data, user_input[b.InputId].(map[string]interface{}))
}
