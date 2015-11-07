package pdf

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
)

type GoFpdf struct {
	fpdf   *gofpdf.Fpdf
	layout Layout
}

// NewGoFpdf generates a new layout based on a Layout interface and uses the
// jung-kurt/gofpdf library to process all the information.
func NewGoFpdf(l Layout) (pdf *GoFpdf) {
	init := &gofpdf.InitType{
		OrientationStr: l.Orientation(),
		UnitStr:        l.Unit(),
		Size:           gofpdf.SizeType{Wd: l.Width(), Ht: l.Height()},
		FontDirStr:     "./fonts",
	}
	fpdf := gofpdf.NewCustom(init)
	fpdf.SetMargins(0, 0, 0)
	fpdf.AddPage()
	fpdf.SetAutoPageBreak(true, 0)

	pdf = new(GoFpdf)
	pdf.fpdf = fpdf
	pdf.layout = l

	for fontName, fontStyles := range l.Fonts() {
		pdf.LoadFonts(fontName, fontStyles)
	}

	return
}

// Layout returns the current layout that is used to render the PDF document.
func (f GoFpdf) Layout() Layout {
	return f.layout
}

// AddFont adds a font to the PDF.
func (f *GoFpdf) LoadFonts(name string, styles map[string]string) {
	for style, file := range styles {
		styleName := ""
		switch style {
		case "bold":
			styleName = "B"
		case "italic":
			styleName = "I"
		case "bold-italic":
			styleName = "BI"
		default:
			styleName = ""
		}
		f.Document().AddFont(name, styleName, file)
	}
}

// Document returns the actual PDF Document instance which is responsible for
// rendering all the blocks on the page.
func (f *GoFpdf) Document() Document {
	return f.fpdf
}

// Bytes sends back a buffer of bytes which can be used to stream to a webpage.
func (f *GoFpdf) Bytes(buffer *bytes.Buffer) []byte {
	err := f.fpdf.Output(buffer)

	if err != nil {
		f.fpdf.SetError(err)
		return nil
	}

	return buffer.Bytes()
}
