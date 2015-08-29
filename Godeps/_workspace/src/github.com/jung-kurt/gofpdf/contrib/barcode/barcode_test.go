package barcode_test

import (
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
	"github.com/jung-kurt/gofpdf/internal/example"
)

func createPdf() (pdf *gofpdf.Fpdf) {
	pdf = gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPage()
	return
}

func ExampleRegister() {
	pdf := createPdf()

	code := "gofpdf"
	fileStr := example.Filename("contrib_barcode_Register")

	bcode, err := code128.Encode(code)

	if err == nil {
		barcode.Register(bcode)
		barcode.Barcode(pdf, code, 15, 15, 100, 10, false)
	}

	err = pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_barcode_Register.pdf
}

func ExampleRegisterCodabar() {
	pdf := createPdf()

	code := "codabar"
	barcode.RegisterCode128(pdf, code)
	barcode.Barcode(pdf, code, 15, 15, 100, 10, false)

	fileStr := example.Filename("contrib_barcode_RegisterCodabar")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_barcode_RegisterCodabar.pdf
}

func ExampleRegisterCode128() {
	pdf := createPdf()

	code := "code128"
	barcode.RegisterCode128(pdf, code)
	barcode.Barcode(pdf, code, 15, 15, 100, 10, false)

	fileStr := example.Filename("contrib_barcode_RegisterCode128")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_barcode_RegisterCode128.pdf
}

func ExampleRegisterCode39() {
	pdf := createPdf()

	code := "CODE39"
	barcode.RegisterCode39(pdf, code, false, true)
	barcode.Barcode(pdf, code, 15, 15, 100, 10, false)

	fileStr := example.Filename("contrib_barcode_RegisterCode39")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_barcode_RegisterCode39.pdf
}

func ExampleRegisterDataMatrix() {
	pdf := createPdf()

	code := "datamatrix"
	barcode.RegisterDataMatrix(pdf, code)
	barcode.Barcode(pdf, code, 15, 15, 20, 20, false)

	fileStr := example.Filename("contrib_barcode_RegisterDataMatrix")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_barcode_RegisterDataMatrix.pdf
}

func ExampleRegisterEAN() {
	pdf := createPdf()

	code := "96385074"
	barcode.RegisterEAN(pdf, code)
	barcode.Barcode(pdf, code, 15, 15, 100, 10, false)

	fileStr := example.Filename("contrib_barcode_RegisterEAN")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_barcode_RegisterEAN.pdf
}

func ExampleRegisterQR() {
	pdf := createPdf()

	code := "qrcode"
	barcode.RegisterQR(pdf, code, qr.H, qr.Unicode)
	barcode.Barcode(pdf, code, 15, 15, 20, 20, false)

	fileStr := example.Filename("contrib_barcode_RegisterQR")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_barcode_RegisterQR.pdf
}

func ExampleRegisterTwoOfFive() {
	pdf := createPdf()

	code := "1234567895"
	barcode.RegisterTwoOfFive(pdf, code, true)
	barcode.Barcode(pdf, code, 15, 15, 100, 20, false)

	fileStr := example.Filename("contrib_barcode_RegisterTwoOfFive")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_barcode_RegisterTwoOfFive.pdf
}
