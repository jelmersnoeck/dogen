package pdf

type Template interface {
	Layout() Layout
	LoadBlocks(user_input interface{})
	Blocks() []Block
}
