package templates

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jelmersnoeck/dogen/renderer/documents"
)

// Image represents an image block. It downloads the image from the given URL
// and puts it on the document on the given position with the given dimension.
type Image struct {
	Url   string  `mapstructure:"url"`
	X     float64 `mapstructure:"x"`
	Y     float64 `mapstructure:"y"`
	W     float64 `mapstructure:"width"`
	H     float64 `mapstructure:"height"`
	Mime  string
	Image io.Reader
}

// Parse puts a HTTP Image on the PDF.
func (i *Image) Parse(doc documents.Document) {
	tp := doc.ImageTypeFromMime(i.Mime)
	doc.RegisterImageReader(i.Url, tp, i.Image)
	doc.Image(i.Url, i.X, i.Y, i.W, i.H, false, tp, 0, "")
}

// Load will look if there is a url key in the user input and if so, it will
// overwrite the given URL from the template with this URL.
func (i *Image) Load(t Template, block_data, user_input map[string]interface{}) {
	if url, ok := user_input["url"]; ok {
		i.Url = url.(string)
	}

	t.WaitGroup().Add(1)
	go i.downloadImage(t)
}

func (i *Image) downloadImage(t Template) {
	defer t.WaitGroup().Done()

	resp, err := http.Get(i.Url)

	if err != nil {
		t.AddError(err)
		return
	}

	defer resp.Body.Close()

	body, readerErr := ioutil.ReadAll(resp.Body)

	if readerErr != nil {
		t.AddError(readerErr)
		return
	}

	i.Image = bytes.NewReader(body)
	i.Mime = resp.Header["Content-Type"][0]
}
