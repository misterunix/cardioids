/*
Copyright 2022 by William Jones
Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
2. A link to the original work. "https://github.com/misterunix/cardioids"
3. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package main

import (
	"fmt"
	"math"
	"os"

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

type ProgVars struct {
	Pid          int
	ibuf0        *gd.Image
	Center       Point
	RadiusColor  gd.Color
	BackGndColor gd.Color
	Edge         []Point
	CircleSteps  int
}

var Program ProgVars

func Cardioid(startpos int, modstep int, hueStart float64) {
	l := len(Program.Edge)

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
		c := Program.ibuf0.ColorAllocateAlpha(int(r), int(g), int(b), 50)
		Program.ibuf0.Line(int(Program.Edge[startIndex].X), int(Program.Edge[startIndex].Y), int(Program.Edge[endIndex].X), int(Program.Edge[endIndex].Y), c)
		startIndex += 1
		startIndex = startIndex % l
		endIndex = (endIndex + modstep) % l
	}
}

func main() {

	fmt.Println("Program cardioid start.")

	Program.Pid = os.Getpid()

	step := 2

	width := 600
	height := 600
	diameter := float64(width) * .95
	radius := diameter / 2.0
	Program.Center.X = float64(width) / 2.0
	Program.Center.Y = float64(height) / 2.0

	Program.ibuf0 = gd.CreateTrueColor(width, height)
	Program.BackGndColor = Program.ibuf0.ColorAllocate(0x00, 0x00, 0x00)
	Program.RadiusColor = Program.ibuf0.ColorAllocateAlpha(0xFF, 0xFF, 0xFF, 0)

	Program.ibuf0.Fill(width/2, height/2, Program.BackGndColor)
	Program.ibuf0.Ellipse(int(Program.Center.X), int(Program.Center.Y), int(diameter), int(diameter), Program.RadiusColor)

	for i := 0; i < 360; i += step {
		p := Point{}
		r := float64(i) * DEG2RAD
		p.X = (math.Cos(r) * radius) + Program.Center.X
		p.Y = (math.Sin(r) * radius) + Program.Center.Y
		Program.Edge = append(Program.Edge, p)
	}

	for loop := 0; loop < 360; loop++ {

		// Clear the background of the image. No transparency.
		Program.ibuf0.Fill(width/2, height/2, Program.BackGndColor)

		// Draw the outside circle.
		Program.ibuf0.Ellipse(int(Program.Center.X), int(Program.Center.Y), int(diameter), int(diameter), Program.RadiusColor)

		// Draw the cardioid
		Cardioid(0, 2, float64(loop))

		/*
			l := len(Program.Edge)
			m := 3
			hss := 360 / float64(l)
			startIndex := 0
			endIndex := startIndex + m
			for i := 0; i < l; i++ {
				hs := math.Abs((hss * float64(i)))
				h := math.Mod(hs, 360)
				r, g, b := hsl.HSLtoRGB(h, 90, 90)
				c := Program.ibuf0.ColorAllocateAlpha(int(r), int(g), int(b), 50)
				Program.ibuf0.Line(int(Program.Edge[startIndex].X), int(Program.Edge[startIndex].Y), int(Program.Edge[endIndex].X), int(Program.Edge[endIndex].Y), c)
				startIndex += 1
				startIndex = startIndex % l
				endIndex = (endIndex + m) % l
			}

			startIndex = l / 3
			endIndex = startIndex + m
			for i := 0; i < l; i++ {
				hs := math.Abs((hss * float64(i)) + 120)
				h := math.Mod(hs, 360)
				r, g, b := hsl.HSLtoRGB(h, 90, 90)
				c := Program.ibuf0.ColorAllocateAlpha(int(r), int(g), int(b), 50)
				Program.ibuf0.Line(int(Program.Edge[startIndex].X), int(Program.Edge[startIndex].Y), int(Program.Edge[endIndex].X), int(Program.Edge[endIndex].Y), c)
				startIndex += 1
				startIndex = startIndex % l
				endIndex = (endIndex + m) % l
			}

			startIndex = l/3 + (l / 3)
			endIndex = startIndex + m
			for i := 0; i < l; i++ {
				hs := math.Abs((hss * float64(i)) + 240)
				h := math.Mod(hs, 360)
				r, g, b := hsl.HSLtoRGB(h, 90, 90)
				c := Program.ibuf0.ColorAllocateAlpha(int(r), int(g), int(b), 50)
				Program.ibuf0.Line(int(Program.Edge[startIndex].X), int(Program.Edge[startIndex].Y), int(Program.Edge[endIndex].X), int(Program.Edge[endIndex].Y), c)
				startIndex += 1
				startIndex = startIndex % l
				endIndex = (endIndex + m) % l
			}
		*/

		//filename := fmt.Sprintf("images/%06d.png", Program.Pid)
		filename := fmt.Sprintf("images/%06d.png", loop)
		Program.ibuf0.Png(filename)
	}
}
