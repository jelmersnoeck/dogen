package pdf_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/renderer/pdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LayoutSuite struct {
	suite.Suite
}

func TestLayoutSuite(t *testing.T) {
	suite.Run(t, new(LayoutSuite))
}

func (s *LayoutSuite) TestNewPageLayoutDefaults() {
	l, err := pdf.NewPageLayout("", "", 1, 1)

	assert.Equal(s.T(), "L", l.Orientation())
	assert.Equal(s.T(), "mm", l.Unit())
	assert.Nil(s.T(), err)
}

func (s *LayoutSuite) TestWidthError() {
	_, err := pdf.NewPageLayout("", "", 0, 1)

	assert.NotNil(s.T(), err)
}

func (s *LayoutSuite) TestHeightError() {
	_, err := pdf.NewPageLayout("", "", 1, 0)

	assert.NotNil(s.T(), err)
}

func (s *LayoutSuite) TestWidth() {
	l, _ := pdf.NewPageLayout("", "", 15, 20)

	assert.EqualValues(s.T(), 15, l.Width())
}

func (s *LayoutSuite) TestHeight() {
	l, _ := pdf.NewPageLayout("", "", 15, 20)

	assert.EqualValues(s.T(), 20, l.Height())
}
