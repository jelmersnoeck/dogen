package templates_test

import (
	"testing"

	"github.com/jelmersnoeck/dogen/renderer/mocks"
	"github.com/jelmersnoeck/dogen/renderer/templates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserInputBlockSuite struct {
	suite.Suite
}

func TestUserInputBlockSuite(t *testing.T) {
	suite.Run(t, new(UserInputBlockSuite))
}

func (s *UserInputBlockSuite) TestParse() {
	input_block := &mocks.Block{}
	user_input := &templates.UserInput{"test", input_block}
	document := &mocks.Document{}

	input_block.On("Parse", document).Return(true)

	user_input.Parse(document)

	input_block.AssertExpectations(s.T())
}

func (s *UserInputBlockSuite) TestParseWithoutBlock() {
	temp := &mocks.Template{}
	ui := &templates.UserInput{}
	ui.InputId = "test-key"

	block_data := map[string]interface{}{"optional": true}
	user_data := map[string]interface{}{"no-test-key": ""}

	ui.Load(temp, block_data, user_data)
}

func (s *UserInputBlockSuite) TestOptionalLoad() {
	temp := &mocks.Template{}
	ui := &templates.UserInput{}
	ui.InputId = "test-key"

	block_data := map[string]interface{}{"optional": true}
	user_data := map[string]interface{}{"no-test-key": ""}

	ui.Load(temp, block_data, user_data)
}

func (s *UserInputBlockSuite) TestMandatoryLoadWithoutKey() {
	temp := &mocks.Template{}
	ui := &templates.UserInput{}
	ui.InputId = "test-key"

	block_data := map[string]interface{}{"optional": false}
	user_data := map[string]interface{}{"no-test-key": ""}

	err := templates.ErrInputIdMandatory("test-key")
	temp.On("AddError", err).Return(true)

	ui.Load(temp, block_data, user_data)
	assert.Nil(s.T(), ui.Block)

	temp.AssertExpectations(s.T())
}

func (s *UserInputBlockSuite) TestLoad() {
	temp := &mocks.Template{}
	block := &mocks.Block{}
	ui := &templates.UserInput{}
	ui.InputId = "test-key"

	block_attributes := map[string]interface{}{"type": "image"}
	block_data := map[string]interface{}{"block_data": block_attributes}

	input_attributes := map[string]interface{}{"url": "my-url"}
	input_data := map[string]interface{}{"test-key": input_attributes}

	temp.On("LoadBlock", block_attributes, input_attributes).Return(block)

	ui.Load(temp, block_data, input_data)

	temp.AssertExpectations(s.T())

	assert.NotNil(s.T(), ui.Block)
}
