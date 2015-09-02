package mocks

import "github.com/jelmersnoeck/noscito/easypdf"
import "github.com/stretchr/testify/mock"

type Block struct {
	mock.Mock
}

func (_m *Block) Parse(pdf *easypdf.EasyPDF, input map[string]interface{}) {
	_m.Called(pdf, input)
}
