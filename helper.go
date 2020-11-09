package colorname

import (
	"encoding/hex"
	"errors"
	"image/color"
)

// ErrInvalidHexFormat is returned if a passed
// color hex string has an invalid format.
var ErrInvalidHexFormat = errors.New("invalid color format")

func toRGB(val int) (r, g, b int) {
	b = val & 255
	g = val >> 8 & 255
	r = val >> 16
	return
}

func toClrRGBA(r, g, b int) *color.RGBA {
	return &color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func abs(v int) int {
	if v < 0 {
		v *= -1
	}
	return v
}

func fromHex(hexVal string) (r, g, b int, err error) {
	if hexVal == "" {
		err = ErrInvalidHexFormat
		return
	}

	if hexVal[0] == '#' {
		return fromHex(hexVal[1:])
	}

	v, err := hex.DecodeString(hexVal)
	if err != nil {
		return
	}

	if len(v) < 3 {
		err = ErrInvalidHexFormat
		return
	}

	r = int(v[0])
	g = int(v[1])
	b = int(v[2])

	return
}
