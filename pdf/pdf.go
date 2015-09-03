// Package PDF contains functionality to draw data onto a PDF and render said
// PDF to a file or string buffer.
package pdf

import (
	"bytes"
	"io"

	"github.com/jung-kurt/gofpdf"
)

// This is basically an interface to wrap gofpdf.Fpdf in.
type PdfDrawer interface {
	AddPage()
	SetError(err error)
	Output(w io.Writer) error
	Error() error
}

type Pdf interface {
	Bytes(buffer *bytes.Buffer) []byte
	Error() error
	Drawer() *gofpdf.Fpdf
	ParseBlocks(blocks []Block)
}

type EasyPdf struct {
	drawer *gofpdf.Fpdf
	Size   Size
}

func NewPdf(size Size) (pdf *EasyPdf) {
	init := &gofpdf.InitType{
		OrientationStr: size.Orientation(),
		UnitStr:        size.Unit(),
		Size:           gofpdf.SizeType{Wd: size.Width(), Ht: size.Height()},
	}
	f := gofpdf.NewCustom(init)
	f.AddPage()

	pdf = new(EasyPdf)
	pdf.drawer = f
	pdf.Size = size
	return
}

func (f *EasyPdf) Error() error {
	return f.Drawer().Error()
}

func (f *EasyPdf) Drawer() *gofpdf.Fpdf {
	return f.drawer
}

func (f *EasyPdf) ParseBlocks(blocks []Block) {
	blocks[0].Parse(f)
}

func (f *EasyPdf) Bytes(buffer *bytes.Buffer) []byte {
	err := f.Drawer().Output(buffer)

	if err != nil {
		f.Drawer().SetError(err)
		return nil
	}

	return buffer.Bytes()
}
