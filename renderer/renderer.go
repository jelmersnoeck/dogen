// Package renderer exists out of multiple parts which, combined, will transform
// data into a specified ouptput format.
package renderer

import (
	"github.com/jelmersnoeck/dogen/renderer/documents"
	"github.com/jelmersnoeck/dogen/renderer/templates"
)

// Render takes a template which it will use to output a document in a specific
// format.
// If successful, an array of bytes will be returned containing the output of
// the document.
// If unsuccessful, an empty array will be returned with a related error.
func Render(template templates.Template, doc documents.Document) {
	for _, block := range template.Blocks() {
		block.Parse(doc)
	}
}
