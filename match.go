package colorname

import "image/color"

// ColorMatch wraps the matched
// color name, RGBA value, the
// RGBA value of the original
// color and the average diff
// between them.
type ColorMatch struct {
	Name    string
	Clr     *color.RGBA
	Actual  *color.RGBA
	AvgDiff float64
}
