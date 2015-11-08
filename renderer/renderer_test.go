package renderer_test

import (
	"bytes"
	"io/ioutil"
	"path"
	"testing"
	"time"

	"github.com/jelmersnoeck/dogen/renderer"
	"github.com/jelmersnoeck/dogen/renderer/documents/pdf"
	"github.com/jelmersnoeck/dogen/renderer/templates"
	"github.com/jelmersnoeck/dogen/renderer/utils"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/suite"
)

type RendererSuite struct {
	suite.Suite
}

func TestRendererSuite(t *testing.T) {
	suite.Run(t, new(RendererSuite))
}

func (s *RendererSuite) TestRender() {
	data := map[string]interface{}{
		"demo_text":  map[string]interface{}{"text": "Jelmer Snoeck - Siphoc"},
		"demo_image": map[string]interface{}{"url": "https://golang.org/doc/gopher/frontpage.png"},
	}

	compareReferencePDF(s, "input-test-pdf", data)
}

func compareReferencePDF(s *RendererSuite, name string, data map[string]interface{}) {
	gofpdf.SetDefaultCatalogSort(true)
	gofpdf.SetDefaultCreationDate(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	envPath, err := utils.EnvironmentPath()

	if err != nil {
		s.T().Errorf(err.Error())
	}

	filepath := path.Join(envPath + "/renderer/templates/fixtures/" + name + ".pdf")
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		s.T().Errorf(err.Error())
	}

	template, _ := templates.LoadJsonTemplate(name, data)

	f := pdf.NewGoFpdf(template.Layout())
	renderer.Render(template, f.Document())

	buffer := bytes.NewBufferString("")

	err = gofpdf.CompareBytes(f.Bytes(buffer), file)

	if err != nil {
		s.T().Errorf(err.Error())
	}
}
