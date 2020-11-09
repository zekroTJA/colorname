// Package colorname finds named colors which match
// or are close to a specified color value.
package colorname

import (
	"image/color"
)

// FindRGB returns an array of matches of named
// colors which are closest to the given RGB
// value by average difference per color.
//
// If two named colors have exacly the same
// average diff to the passed color values,
// multiple values may be contained in the
// result array.
//
// The result array will always contain at
// least one match.
func FindRGB(r, g, b int) (matches []*ColorMatch) {
	matches = make([]*ColorMatch, 0)
	minAvgDiff := float64(^uint(0) >> 1)

	actual := toClrRGBA(r, g, b)

	for name, numVal := range colorNames {
		vR, vG, vB := toRGB(numVal)
		avgDiff := float64(abs(r-vR)+abs(g-vG)+abs(b-vB)) / 3

		if avgDiff > minAvgDiff {
			continue
		}

		match := &ColorMatch{
			Name:    name,
			Actual:  actual,
			Clr:     toClrRGBA(vR, vG, vB),
			AvgDiff: avgDiff,
		}

		if avgDiff == minAvgDiff {
			matches = append(matches, match)
		} else {
			minAvgDiff = avgDiff
			matches = []*ColorMatch{match}
		}
	}

	return
}

// FindRGBA is shorthand for FindRGB using a
// color.RGBA object.
func FindRGBA(rgba *color.RGBA) []*ColorMatch {
	return FindRGB(int(rgba.R), int(rgba.G), int(rgba.B))
}

// FindNum is shorthand for FindRGB using an
// integer color value.
func FindNum(num int) []*ColorMatch {
	return FindRGB(toRGB(num))
}

// FindHexStr is shorthand for FindRGB using a
// hey string as color value.
//
// An error is returned, if the passed string
// has an invalid format.
func FindHexStr(hexVal string) ([]*ColorMatch, error) {
	r, g, b, err := fromHex(hexVal)
	if err != nil {
		return nil, err
	}

	return FindRGB(r, g, b), nil
}
