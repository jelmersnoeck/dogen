package mocks

import "github.com/jelmersnoeck/noscito/services/easypdf"
import "github.com/stretchr/testify/mock"

type Block struct {
	mock.Mock
}

func (_m *Block) Parse(pdf *easypdf.EasyPDF, input map[string]interface{}) {
	_m.Called(pdf, input)
}
func (_m *Block) Load(data interface{}) {
	_m.Called(data)
}
