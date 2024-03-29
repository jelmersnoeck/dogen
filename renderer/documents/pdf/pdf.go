// Package PDF contains functionality to draw data onto a PDF and render said
// PDF to a file or string buffer.
package pdf

import (
	"bytes"

	"github.com/jelmersnoeck/dogen/renderer/documents"
	"github.com/jelmersnoeck/dogen/renderer/layouts"
)

// Pdf houses all the wrapper functionality about our actual PDF instance. It
// has a Document which it uses to wrypte all the data to and render the pages.
// The Layout represents the page layout.
type Pdf interface {
	Bytes(buffer *bytes.Buffer) []byte
	LoadFonts(name string, styles map[string]string)
	Layout() layouts.Layout
	Document() documents.Document
}
