package templates

import (
	"errors"

	"github.com/jelmersnoeck/dogen/renderer/documents"
)

var ErrInputIdMandatory = func(input_id string) error {
	return errors.New("Input field " + input_id + " required but not present")
}

// UserInput takes a Block and InputId which it will use on loading the template
// to match with the UserInput given through the API.
type UserInput struct {
	InputId string `mapstructure:"input_id"`
	Block   Block
}

// Parse will recursively parse the blocks that are under this UserInput.
func (b *UserInput) Parse(doc documents.Document) {
	if b.Block != nil {
		b.Block.Parse(doc)
	}
}

// Load will recursively load blocks that fall under this block. It will select
// the correct input from the user_input based on the input_id that is specified
// in the block_data.
//
// If the data is optional, it will see if there is an input key that matches
// the InputId. If not, then it will skip adding the block. If it is not
// optional then it will add an error to the template if the key is not present.
func (b *UserInput) Load(t Template, block_data, user_input map[string]interface{}) {
	if optional, ok := block_data["optional"]; ok {
		_, present := user_input[b.InputId]

		if !present {
			if !optional.(bool) {
				t.AddError(ErrInputIdMandatory(b.InputId))
			}

			return
		}
	}

	b.Block = t.LoadBlock(
		block_data["block_data"].(map[string]interface{}),
		user_input[b.InputId].(map[string]interface{}),
	)
}
