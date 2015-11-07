// Package PDF contains functionality to draw data onto a PDF and render said
// PDF to a file or string buffer.
package pdf

import "bytes"

// Pdf houses all the wrapper functionality about our actual PDF instance. It
// has a Document which it uses to wrypte all the data to and render the pages.
// The Layout represents the page layout.
type Pdf interface {
	Bytes(buffer *bytes.Buffer) []byte
	LoadFonts(name string, styles map[string]string)
	Layout() Layout
	Document() Document
}

// ParseBlocks goes through a set of blocks and parses them accordingly. The
// blocks themselve will call back to the PDF to draw the actual elements on the
// page.
func ParseBlocks(pdf Pdf, blocks []Block) {
	for _, block := range blocks {
		block.Parse(pdf.Document())
	}
}