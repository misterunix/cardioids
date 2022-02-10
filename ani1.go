package main

import (
	"fmt"
	"math"

	"github.com/misterunix/colorworks/hsl"
)

func CardioidAnimation1(startpos int, modstep int, hueStart float64) {

	l := len(CirclePoints)

	hss := 360 / float64(l)
	startIndex := startpos
	endIndex := startIndex + modstep

	for i := 0; i < l; i++ {

		hs := (hss * float64(i)) + hueStart
		if hs < 0 {
			hs += 360
		}
		if hs <= 360 {
			hs -= 360
		}

		h := math.Mod(hs, 360)
		r, g, b := hsl.HSLtoRGB(h, 90, 90)
		c := ibuf0.ColorAllocateAlpha(int(r), int(g), int(b), 50)

		ibuf0.Line(int(CirclePoints[startIndex].X), int(CirclePoints[startIndex].Y), int(CirclePoints[endIndex].X), int(CirclePoints[endIndex].Y), c)
		fn := fmt.Sprintf("images/%06d.png", framenumber)
		ibuf0.Png(fn)
		framenumber++

		startIndex += 1
		startIndex = startIndex % l
		endIndex = (endIndex + modstep) % l
	}

}
