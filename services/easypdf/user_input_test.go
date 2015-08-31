package easypdf_test

import (
	"encoding/json"
	"testing"

	"github.com/jelmersnoeck/noscito/services/easypdf"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestSuccessfulLoad(c *C) {
	data := make(map[string]interface{})

	byt := []byte(`{
		"input_id": "test",
		"block": {
			"type": "image"
		}
	}`)

	json.Unmarshal(byt, &data)

	ui := easypdf.UserInput{}
	ui.Load(data)

	c.Assert(ui.Block, FitsTypeOf, &easypdf.Image{})
	c.Assert(ui.InputId, Equals, "test")
}
