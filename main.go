package main

import (
	"fmt"
	"math"

	gd "github.com/misterunix/cgo-gd"
	"github.com/misterunix/colorworks/hsl"
)

const (
	DEG2RAD = 0.0174532925
	RAD2DEG = 57.2957795130
)

type Point struct {
	X float64
	Y float64
}

var Edge []Point
var Center Point

var ibuf0 *gd.Image

func main() {

	fmt.Println("Program cardioid start.")

	step := 2

	width := 600
	height := 600
	diameter := float64(width) * .95
	radius := diameter / 2.0
	Center.X = float64(width) / 2.0
	Center.Y = float64(height) / 2.0

	ibuf0 = gd.CreateTrueColor(width, height)
	c1 := ibuf0.ColorAllocateAlpha(0xFF, 0xFF, 0xFF, 0)
	ibuf0.Ellipse(int(Center.X), int(Center.Y), int(diameter), int(diameter), c1)

	for i := 0; i < 360; i += step {
		p := Point{}
		r := float64(i) * DEG2RAD
		p.X = (math.Cos(r) * radius) + Center.X
		p.Y = (math.Sin(r) * radius) + Center.Y
		Edge = append(Edge, p)
	}

	l := len(Edge)
	m := 3
	hss := 360 / float64(l)
	startIndex := 0
	endIndex := startIndex + m
	for i := 0; i < l; i++ {
		hs := math.Abs((hss * float64(i)) - 270)
		h := math.Mod(hs, 360)
		r, g, b := hsl.HSLtoRGB(h, 90, 90)
		c := ibuf0.ColorAllocateAlpha(int(r), int(g), int(b), 0)
		ibuf0.Line(int(Edge[startIndex].X), int(Edge[startIndex].Y), int(Edge[endIndex].X), int(Edge[endIndex].Y), c)
		startIndex += 1
		startIndex = startIndex % l
		endIndex = (endIndex + m) % l
	}

	startIndex = l / 3
	endIndex = startIndex + m
	for i := 0; i < l; i++ {
		hs := math.Abs((hss * float64(i)) - 270)
		h := math.Mod(hs, 360)
		r, g, b := hsl.HSLtoRGB(h, 90, 90)
		c := ibuf0.ColorAllocateAlpha(int(r), int(g), int(b), 0)
		ibuf0.Line(int(Edge[startIndex].X), int(Edge[startIndex].Y), int(Edge[endIndex].X), int(Edge[endIndex].Y), c)
		startIndex += 1
		startIndex = startIndex % l
		endIndex = (endIndex + m) % l
	}

	/*
		startIndex = l / 3
		endIndex = startIndex + m
		for i := 0; i < l; i++ {
			ibuf0.Line(int(Edge[startIndex].X), int(Edge[startIndex].Y),
				int(Edge[endIndex].X), int(Edge[endIndex].Y), c1)
			startIndex += 1
			startIndex = startIndex % l
			endIndex = (endIndex + m) % l
		}

		startIndex = l / 6
		endIndex = startIndex + m
		for i := 0; i < l; i++ {
			ibuf0.Line(int(Edge[startIndex].X), int(Edge[startIndex].Y),
				int(Edge[endIndex].X), int(Edge[endIndex].Y), c1)
			startIndex += 1
			startIndex = startIndex % l
			endIndex = (endIndex + m) % l
		}
	*/
	//	for _, p := range Edge {
	//		ibuf0.Line(int(Center.X), int(Center.Y), int(p.X), int(p.Y), c1)
	//	}

	ibuf0.Png("images/test.png")

}