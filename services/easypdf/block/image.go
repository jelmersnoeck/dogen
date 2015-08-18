package block

import (
	"github.com/jung-kurt/gofpdf"
	"net/http"
)

type Image struct {
	Url        string
	X, Y, W, H float64
}

func (i *Image) Parse(pdf *gofpdf.Fpdf) {
	registerRemoteImage(pdf, i.Url)
	pdf.Image(i.Url, i.X, i.Y, i.W, i.H, false, "", 0, "")
}

func registerRemoteImage(f *gofpdf.Fpdf, urlStr string) (info *gofpdf.ImageInfoType) {
	resp, err := http.Get(urlStr)

	if err != nil {
		f.SetError(err)
		return
	}

	defer resp.Body.Close()

	tp := f.ImageTypeFromMime(resp.Header["Content-Type"][0])

	return f.RegisterImageReader(urlStr, tp, resp.Body)
}
