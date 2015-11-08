package pdf_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PdfSuite struct {
	suite.Suite
}

func TestPdfSuite(t *testing.T) {
	suite.Run(t, new(PdfSuite))
}
