package pdf

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

type GoFpdf struct {
	fpdf *gofpdf.Fpdf
	Size Size
}

func NewGoFpdf(size Size) (pdf *GoFpdf) {
	init := &gofpdf.InitType{
		OrientationStr: size.Orientation(),
		UnitStr:        size.Unit(),
		Size:           gofpdf.SizeType{Wd: size.Width(), Ht: size.Height()},
	}
	fpdf := gofpdf.NewCustom(init)
	fpdf.AddPage()

	pdf = new(GoFpdf)
	pdf.fpdf = fpdf
	pdf.Size = size
	return
}

func (f *GoFpdf) HttpImage(url string, x, y, w, h float64, flow bool, tp string, link int, linkStr string) {
	httpimg.Register(f.fpdf, url, "")
	f.fpdf.Image(url, x, y, w, h, flow, tp, link, linkStr)
}

func (f *GoFpdf) ParseBlocks(blocks []Block) {
	for _, block := range blocks {
		block.Parse(f)
	}
}

func (f *GoFpdf) Bytes(buffer *bytes.Buffer) []byte {
	err := f.fpdf.Output(buffer)

	if err != nil {
		f.fpdf.SetError(err)
		return nil
	}

	return buffer.Bytes()
}
