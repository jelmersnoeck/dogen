package utils

import "strconv"

// HexToRGB converts an Hex string to a RGB triple.
func HexToRGB(h string) (int, int, int) {
	if len(h) > 0 && h[0] == '#' {
		h = h[1:]
	}
	if len(h) == 3 {
		h = h[:1] + h[:1] + h[1:2] + h[1:2] + h[2:] + h[2:]
	}
	if len(h) == 6 {
		if rgb, err := strconv.ParseUint(string(h), 16, 32); err == nil {
			return int(rgb >> 16), int((rgb >> 8) & 0xFF), int(rgb & 0xFF)
		}
	}
	return 0, 0, 0
}
