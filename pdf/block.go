package pdf

// Block represents an item in a template file which needs to be drawn onto the
// PDF. We will first Load the block through the template and assign all correct
// UserInput to it and then we'll Parse the block onto a Document.
type Block interface {
	Parse(doc Document)
	Load(t Template, block_data, user_input map[string]interface{})
}
