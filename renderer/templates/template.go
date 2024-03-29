// Package templates provides a templating system to draw items on your
// documents.
package templates

import (
	"sync"

	"github.com/jelmersnoeck/dogen/renderer/layouts"
)

// Template is what we use to combine what a user has drawn in the PDF generator
// to use as a template with the UserInput they provide through the API.
type Template interface {
	AddError(err error)
	Blocks() []Block
	Errors() []error
	Layout() layouts.Layout
	LoadBlock(raw_block, raw_input map[string]interface{}) Block
	LoadBlocks(user_input interface{})
	WaitGroup() *sync.WaitGroup
}
