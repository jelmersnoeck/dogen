package pdf

type Block interface {
	Parse(doc Document)
	Load(t Template, block_data, user_input map[string]interface{})
}
