package pdf

type Template interface {
	Layout() Layout
	Blocks() []Block
}
