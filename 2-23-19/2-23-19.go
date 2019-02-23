package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

const width = 1000
const height = 1000
const circleRadius = 500.0

var namer = fn.New()

func pickPointOnCircle(rand *rand.Rand, radius float64) (float64, float64) {
	a := rand.Float64() * 2 * math.Pi
	x := radius * math.Cos(a)
	y := radius * math.Sin(a)

	return x, y
}

func drawRipple(dc *gg.Context, rand *rand.Rand, cx, cy float64) {
	for r := 1.0; r < width; r = math.Max(r+10, r*1.2) {
		for i := 0; i < int(r*4); i++ {
			x, y := pickPointOnCircle(rand, float64(r))
			dc.DrawPoint(x+cx, y+cy, 1)
			dc.Stroke()
		}
	}
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(1)

	drawRipple(dc, r1, 0, 500)
	drawRipple(dc, r1, 500, 0)
	drawRipple(dc, r1, 500, 1000)
	drawRipple(dc, r1, 1000, 500)

	// dc.DrawCircle(0, 0, 500)

	dc.SavePNG(namer.NameWithFileType("png"))
}
