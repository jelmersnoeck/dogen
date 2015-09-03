// Package PDF contains functionality to draw data onto a PDF and render said
// PDF to a file or string buffer.
package pdf

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
)

// This is basically an interface to wrap gofpdf.Fpdf in.
type PdfDrawer interface {
	AddPage()
}

type Pdf interface {
	Drawer() *PdfDrawer
	Bytes() []byte
}

type EasyPdf struct {
	drawer PdfDrawer
}

func NewPdf() (pdf *EasyPdf) {
	f := gofpdf.New("P", "mm", "A4", "")
	f.AddPage()

	pdf = &EasyPdf{f}
	return
}

func (f *EasyPdf) Drawer() *PdfDrawer {
	return &f.drawer
}

func (f *EasyPdf) Bytes() []byte {
	buffer := bytes.NewBufferString("")
	return buffer.Bytes()
}
