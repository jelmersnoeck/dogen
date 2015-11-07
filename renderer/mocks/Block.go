package mocks

import "github.com/jelmersnoeck/dogen/renderer/pdf"
import "github.com/stretchr/testify/mock"

type Block struct {
	mock.Mock
}

func (_m *Block) Parse(doc pdf.Document) {
	_m.Called(doc)
}
func (_m *Block) Load(t pdf.Template, block_data map[string]interface{}, user_input map[string]interface{}) {
	_m.Called(t, block_data, user_input)
}
