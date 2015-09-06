package pdf

type Block interface {
	Parse(pdf Pdf)
	Load(t Template, block_data, user_input map[string]interface{})
}
