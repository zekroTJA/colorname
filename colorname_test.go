package colorname

import (
	"image/color"
	"testing"
)

func TestFindRGB_Exact(t *testing.T) {
	const r, g, b = 106, 13, 173
	matches := FindRGB(r, g, b)
	if len(matches) == 0 {
		t.Error("matches length was 0")
	}

	const expectedName = "Purple"
	match := matches[0]
	if match.Name != expectedName {
		t.Errorf("match name was %s instead of %s", match.Name, expectedName)
	}
	if match.Actual.R != r || match.Actual.G != g || match.Actual.B != b {
		t.Errorf("actual color was (R %d, G %d, B %d) instead of (R %d, G %d, B %d)",
			match.Actual.R, match.Actual.G, match.Actual.B, r, g, b)
	}
	if match.AvgDiff != 0 {
		t.Errorf("diff was %f instead of 0", match.AvgDiff)
	}
}

func TestFindRGB_Close(t *testing.T) {
	const r, g, b = 106 + 2, 13 - 6, 173 + 3
	const diff = float64(2+6+3) / 3
	matches := FindRGB(r, g, b)
	if len(matches) == 0 {
		t.Error("matches length was 0")
	}

	const expectedName = "Purple"
	match := matches[0]
	if match.Name != expectedName {
		t.Errorf("match name was %s instead of %s", match.Name, expectedName)
	}
	if match.Actual.R != r || match.Actual.G != g || match.Actual.B != b {
		t.Errorf("actual color was (R %d, G %d, B %d) instead of (R %d, G %d, B %d)",
			match.Actual.R, match.Actual.G, match.Actual.B, r, g, b)
	}
	if match.AvgDiff != diff {
		t.Errorf("diff was %f instead of %f", match.AvgDiff, diff)
	}
}

func TestFindRGBA(t *testing.T) {
	const r, g, b = 106, 13, 173
	matches := FindRGBA(&color.RGBA{r, g, b, 255})

	const expectedName = "Purple"
	match := matches[0]
	if match.Name != expectedName {
		t.Errorf("match name was %s instead of %s", match.Name, expectedName)
	}
	if match.Actual.R != r || match.Actual.G != g || match.Actual.B != b {
		t.Errorf("actual color was (R %d, G %d, B %d) instead of (R %d, G %d, B %d)",
			match.Actual.R, match.Actual.G, match.Actual.B, r, g, b)
	}
	if match.AvgDiff != 0 {
		t.Errorf("diff was %f instead of 0", match.AvgDiff)
	}
}
func TestFindNum(t *testing.T) {
	const r, g, b = 106, 13, 173
	const numVal = 0x6a0dad
	matches := FindNum(numVal)

	const expectedName = "Purple"
	match := matches[0]
	if match.Name != expectedName {
		t.Errorf("match name was %s instead of %s", match.Name, expectedName)
	}
	if match.Actual.R != r || match.Actual.G != g || match.Actual.B != b {
		t.Errorf("actual color was (R %d, G %d, B %d) instead of (R %d, G %d, B %d)",
			match.Actual.R, match.Actual.G, match.Actual.B, r, g, b)
	}
	if match.AvgDiff != 0 {
		t.Errorf("diff was %f instead of 0", match.AvgDiff)
	}
}

func TestFindHexStr(t *testing.T) {
	_, err := FindHexStr("")
	if err == nil {
		t.Error("did not return error on empty string")
	}

	_, err = FindHexStr("#1111")
	if err == nil {
		t.Error("did not return error on invalid format")
	}

	_, err = FindHexStr("zzzzzz")
	if err == nil {
		t.Error("did not return error on invalid format")
	}

	const r, g, b = 106, 13, 173
	const hexVal = "#6a0dad"
	matches, err := FindHexStr(hexVal)
	if err != nil {
		t.Error(err)
	}

	const expectedName = "Purple"
	match := matches[0]
	if match.Name != expectedName {
		t.Errorf("match name was %s instead of %s", match.Name, expectedName)
	}
	if match.Actual.R != r || match.Actual.G != g || match.Actual.B != b {
		t.Errorf("actual color was (R %d, G %d, B %d) instead of (R %d, G %d, B %d)",
			match.Actual.R, match.Actual.G, match.Actual.B, r, g, b)
	}
	if match.AvgDiff != 0 {
		t.Errorf("diff was %f instead of 0", match.AvgDiff)
	}
}

// --- BENCHMARKS ------------------------------------------------

func BenchmarkFindRGB(b *testing.B) {
	const cr, cg, cb = 106, 13, 173
	for n := 0; n < b.N; n++ {
		FindRGB(cr, cg, cb)
	}
}
