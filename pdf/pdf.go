// Package PDF contains functionality to draw data onto a PDF and render said
// PDF to a file or string buffer.
package pdf

import "bytes"

type Pdf interface {
	HttpImage(url string, x, y, w, h float64, flow bool, tp string, link int, linkStr string)
	Position() (float64, float64)
	Bytes(buffer *bytes.Buffer) []byte
}
