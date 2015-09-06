package pdf

type Template interface {
	AddError(err error)
	Blocks() []Block
	Errors() []error
	Layout() Layout
	LoadBlock(raw_block, raw_input map[string]interface{}) Block
	LoadBlocks(user_input interface{})
}
