package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zekroTJA/colorname"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	r, g, b := rand.Intn(256), rand.Intn(256), rand.Intn(256)

	fmt.Printf("Random Color: R %d, G %d, B %d\n\nMatched Colors:\n\n", r, g, b)

	matches := colorname.FindRGB(r, g, b)

	for _, match := range matches {
		fmt.Printf(
			"Name:    %s\n"+
				"RGB:     R %d, G %d, B %d\n"+
				"AvgDiff: %f\n\n",
			match.Name, match.Clr.R, match.Clr.G,
			match.Clr.B, match.AvgDiff)
	}
}
