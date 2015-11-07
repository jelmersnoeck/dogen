package utils_test

import (
	"testing"

	"github.com/jelmersnoeck/dogen/renderer/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TemplateUtilSuite struct {
	suite.Suite
}

func TestTemplateUtilSuite(t *testing.T) {
	suite.Run(t, new(TemplateUtilSuite))
}

func (s *TemplateUtilSuite) TestLoadExistingTemplate() {
	template, err := utils.LoadTemplate("test-pdf")

	assert.Nil(s.T(), err)

	var byt []byte
	assert.IsType(s.T(), byt, template)
}

func (s *TemplateUtilSuite) TestLoadNonExistingTemplate() {
	template, err := utils.LoadTemplate("non-existing")

	assert.Equal(s.T(), utils.NonExistingTemplateError, err)
	assert.Nil(s.T(), template)
}
