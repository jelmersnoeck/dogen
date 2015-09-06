// Package PDF contains functionality to draw data onto a PDF and render said
// PDF to a file or string buffer.
package pdf

import (
	"errors"
)

type Layout interface {
	Orientation() string
	Unit() string
	Width() float64
	Height() float64
}

type PageLayout struct {
	orientation string
	unit        string
	width       float64
	height      float64
}

// NewPageLayout creates a new PageLayout that will be used to create a PDF.
//
// If orientation is empty, the default will be used, which is "L" for landscape.
// If unit is empty, the default will be used, which is "mm" for milimeters.
// If either the width or height are equal or smaller than zero, an error will
// be returned
func NewPageLayout(orientation, unit string, width, height float64) (s *PageLayout, e error) {
	if orientation == "" {
		orientation = "L"
	}

	if unit == "" {
		unit = "mm"
	}

	if width <= 0 {
		e = errors.New("The width of the page needs to be bigger than 0.")
		return
	}

	if height <= 0 {
		e = errors.New("The height of the page needs to be bigger than 0.")
		return
	}

	s = &PageLayout{orientation, unit, width, height}
	return
}

// Orientation gives back the orientation that is used for a page. This can
// either be `L` for landscape or `P` for Portrait.
func (s *PageLayout) Orientation() string {
	return s.orientation
}

// Unit gives back the unit that all the measurements are based on.
func (s *PageLayout) Unit() string {
	return s.unit
}

// Width gives back the width of the page in the given unit.
func (s *PageLayout) Width() float64 {
	return s.width
}

// Height gives back the height of the page in the given unit.
func (s *PageLayout) Height() float64 {
	return s.height
}
