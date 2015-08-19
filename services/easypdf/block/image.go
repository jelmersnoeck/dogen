package block

import (
	"github.com/jung-kurt/gofpdf"
	"io"
	"net/http"
)

type Image struct {
	InputId    string `json:"input_id"`
	Url        string `json:"url"`
	Data       io.Reader
	X, Y, W, H float64
}

func (i *Image) Parse(pdf *gofpdf.Fpdf, data map[string]interface{}) {
	if i.InputId != "" {
		url, ok := data[i.InputId].(string)
		if ok {
			i.Url = url
		}
	}
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
