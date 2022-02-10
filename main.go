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

var CirclePoints []Point

type ProgVars struct {
	Pid          int
	ibuf0        *gd.Image
	Center       Point
	RadiusColor  gd.Color
	BackGndColor gd.Color
	Edge         []Point
	CircleSteps  int
}

var pid int
var ibuf0 *gd.Image
var imageWidth int
var imageHeight int
var numberOfPoints float64
var radius float64
var center Point
var diameter float64
var bkGndColor gd.Color
var circumfrenceColor gd.Color

var framenumber int

// MakePointsAroundCircle : Create slice of points around the circle.
// Needs numberOfPoints and radius
func MakePointsAroundCircle() {
	inc := 360 / numberOfPoints
	for i := 0.0; i < 360.0; i += inc {
		p := Point{}
		r := float64(i) * DEG2RAD
		p.X = (math.Cos(r) * radius) + center.X
		p.Y = (math.Sin(r) * radius) + center.Y
		CirclePoints = append(CirclePoints, p)
	}
}

func Cardioid(startpos int, modstep int, hueStart float64) {
	l := len(CirclePoints)
	fmt.Println(l)

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
		fmt.Printf("i:%d startIndex:%d endIndex:%d\n", i, startIndex, endIndex)
		ibuf0.Line(int(CirclePoints[startIndex].X), int(CirclePoints[startIndex].Y), int(CirclePoints[endIndex].X), int(CirclePoints[endIndex].Y), c)
		startIndex += 1
		startIndex = startIndex % l
		endIndex = (endIndex + modstep) % l
	}
}

func main() {

	fmt.Println("Program cardioid start.")

	pid = os.Getpid()

	//step := 2

	imageWidth = 600
	imageHeight = 600

	diameter = float64(imageWidth) * .95
	radius = diameter / 2.0
	center.X = float64(imageWidth) / 2.0
	center.Y = float64(imageHeight) / 2.0

	ibuf0 = gd.CreateTrueColor(imageWidth, imageHeight)
	bkGndColor = ibuf0.ColorAllocateAlpha(0x00, 0x00, 0x00, 0)
	circumfrenceColor = ibuf0.ColorAllocateAlpha(0xFF, 0xFF, 0xFF, 0)
	framenumber = 0
	//ibuf0.Fill(int(center.X), int(center.Y), bkGndColor)
	//ibuf0.Ellipse(int(center.X), int(center.Y), int(diameter), int(diameter), circumfrenceColor)

	jump := 180
	//for {
	for j := 2; j < 10; j++ {
		//ibuf0.Fill(int(center.X), int(center.Y), bkGndColor)
		ibuf0.FilledRectangle(0, 0, imageWidth, imageHeight, bkGndColor)
		numberOfPoints = float64(jump)
		CirclePoints = CirclePoints[:0]
		MakePointsAroundCircle()
		//Cardioid(0, 6, 0)
		CardioidAnimation1(0, j, ((float64(j)-1)/9.0)*360.0)
		for k := 0; k < 120; k++ {
			fn := fmt.Sprintf("images/%06d.png", framenumber)
			ibuf0.Png(fn)
			framenumber++
		}
	}
	//fn := fmt.Sprintf("images/%f.png", numberOfPoints)
	//ibuf0.Png(fn)
	//jump = jump + 9
	//	if jump >= 360.0 {
	//		break
	//	}
	//}
	//	CircleMod := 360

	/*
		for i := 0; i < 360; i++ {
			p := Point{}
			r := float64(i) * DEG2RAD
			p.X = (math.Cos(r) * radius) + center.X
			p.Y = (math.Sin(r) * radius) + center.Y
			CirclePoints = append(CirclePoints, p)
		}
	*/

	/*
		frame := 0

		mm := 2

		for loop := 0; loop < CircleMod; loop += mm {

			// Clear the background of the image. No transparency.
			ibuf0.Fill(int(center.X), int(center.Y), bkGndColor)

			// Draw the outside circle.
			ibuf0.Ellipse(int(center.X), int(center.Y), int(diameter), int(diameter), circumfrenceColor)

			l := len(CirclePoints)

			// Draw the cardioid

			Cardioid(0, mm, float64(loop))
			Cardioid(l/4, mm, float64(loop))
			Cardioid((l/4)+(l/4), mm, float64(loop))
			Cardioid((l/4)+(l/4)+(l/4), mm, float64(loop))

			//Cardioid(l/3, mm, float64(loop))
			//Cardioid((l/3)+(l/3), mm, float64(loop))

			//filename := fmt.Sprintf("images/%06d.png", Program.Pid)
			filename := fmt.Sprintf("images/%06d.png", frame)
			ibuf0.Png(filename)
			frame++
		}
	*/
}
