package pdf

type Template interface {
	Layout() Layout
	LoadBlock(raw_block, raw_input map[string]interface{}) Block
	LoadBlocks(user_input interface{})
	Blocks() []Block
}
