package pdf

type Block interface {
	Parse(pdf Pdf)
}
