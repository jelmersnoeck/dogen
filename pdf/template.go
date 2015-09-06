package pdf

type Template interface {
	Layout() Layout
	LoadBlock(raw_block, raw_input map[string]interface{})
	LoadBlocks(user_input interface{})
	Blocks() []Block
}
