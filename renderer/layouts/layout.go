// Package layouts represents the layout of an item to be rendered.
package layouts

import (
	"errors"
)

// Layout is used to create a document based on the measurments and unit given.
type Layout interface {
	Orientation() string
	Unit() string
	Width() float64
	Height() float64
	SetFonts(fonts map[string]map[string]string)
	Fonts() map[string]map[string]string
}

// PageLayout represents the Layout that we will use for a normal page in a PDF
// document.
type PageLayout struct {
	orientation string
	unit        string
	width       float64
	height      float64
	fonts       map[string]map[string]string
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

	s = &PageLayout{
		orientation: orientation,
		unit:        unit,
		width:       width,
		height:      height,
	}
	return
}

// Orientation gives back the orientation that is used for a page. This can
// either be `L` for landscape or `P` for Portrait.
func (s *PageLayout) Orientation() string {
	return s.orientation
}

// SetFonts stores the fonts that come from a nested map on the layout.
func (s *PageLayout) SetFonts(fonts map[string]map[string]string) {
	s.fonts = fonts
}

// Fonts returns the array of fonts we'll use in the template.
func (s *PageLayout) Fonts() map[string]map[string]string {
	return s.fonts
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
