package utils_test

import (
	"testing"

	"github.com/jelmersnoeck/noscito/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilsColorSuite struct {
	suite.Suite
}

func TestUtilsColorSuite(t *testing.T) {
	suite.Run(t, new(UtilsColorSuite))
}

func (s *UtilsColorSuite) TestHexToRGBWhite() {
	r, g, b := utils.HexToRGB("ffffff")

	assert.Equal(s.T(), 255, r)
	assert.Equal(s.T(), 255, g)
	assert.Equal(s.T(), 255, b)
}

func (s *UtilsColorSuite) TestHexToRGBBlack() {
	r, g, b := utils.HexToRGB("000000")

	assert.Equal(s.T(), 0, r)
	assert.Equal(s.T(), 0, g)
	assert.Equal(s.T(), 0, b)
}

func (s *UtilsColorSuite) TestHexToRGBBlue() {
	r, g, b := utils.HexToRGB("0000ff")

	assert.Equal(s.T(), 0, r)
	assert.Equal(s.T(), 0, g)
	assert.Equal(s.T(), 255, b)
}
