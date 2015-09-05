package mocks

import "github.com/jelmersnoeck/noscito/pdf"
import "github.com/stretchr/testify/mock"

type Block struct {
	mock.Mock
}

func (_m *Block) Parse(pdf pdf.Pdf) {
	_m.Called(pdf)
}
