package main

import (
	"math"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

var namer = fn.New()

func mapPoint(x1, x2, y1, y2, xt float64) float64 {
	return y1 + (xt-x1)*(y2-y1)/(x2-x1)
}

func drawCircleRing(dc *gg.Context, r, n float64) float64 {
	da := 2 * math.Pi / n
	individR := r * math.Sin(da/2)
	offset := n / 2
	for a := 0.0; a < 2*math.Pi; a += da {
		x := r * math.Cos(a+offset)
		y := r * math.Sin(a+offset)
		dc.DrawCircle(x, y, individR)
		dc.DrawCircle(x, y, individR/2)
		if individR > 2 {
			dc.DrawPoint(x, y, individR/10)
		}
		dc.Stroke()
	}
	return individR
}

func main() {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(1)
	dc.Translate(width/2, height/2)

	r := 17.
	for i := 5.; i < 50; i++ {
		r += 2.1 * drawCircleRing(dc, r, 2*i)
	}

	dc.Stroke()
	// dc.Translate(-minX, -minY)

	dc.SavePNG(namer.NameWithFileType("png"))
}
