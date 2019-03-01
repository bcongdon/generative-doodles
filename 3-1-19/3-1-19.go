package main

import (
	"math"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

const width = 1500
const height = 1500

var namer = fn.New()

func lerp(x1, x2, y1, y2, xt float64) float64 {
	return y1 + (xt-x1)*(y2-y1)/(x2-x1)
}

func drawRect(dc *gg.Context, x1, y1, x2, y2, x3, y3, x4, y4 float64, n int) {
	dc.DrawLine(x1, y1, x2, y2)
	dc.DrawLine(x2, y2, x3, y3)
	dc.DrawLine(x3, y3, x4, y4)
	dc.DrawLine(x4, y4, x1, y1)

	for i := 1; i < n; i++ {
		t := float64(i) / float64(n)
		lx1 := x1 + t*(x2-x1)
		ly1 := lerp(x1, x2, y1, y2, lx1)
		lx2 := x4 + t*(x3-x4)
		ly2 := lerp(x4, x3, y4, y3, lx2)
		dc.DrawLine(lx1, ly1, lx2, ly2)
	}
}

func drawCube(dc *gg.Context, x, y, s float64) {
	s45 := math.Sin(gg.Radians(45))
	c45 := math.Cos(gg.Radians(45))
	drawRect(dc, x, y, x+s*c45, y+s*s45, x+s*c45, y+s+s*s45, x, y+s, int(s/2))
	drawRect(dc, x, y, x+s*c45, y-s*s45, x+math.Sqrt(2*s*s), y, x+s*c45, y+s*s45, int(s/5))
	drawRect(dc, x+s*c45, y+s*s45, x+math.Sqrt(2*s*s), y, x+math.Sqrt(2*s*s), y+s, x+s*c45, y+s+s*s45, int(s/8))
}

func main() {

	dc := gg.NewContext(width, height)

	dc.SetRGB(240./255, 240./255, 240./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.Scale(1.1, 1)

	s := 85.
	for xIdx := -1; math.Ceil(float64(xIdx)*math.Sqrt(2*s*s)) < width; xIdx++ {
		x := float64(xIdx) * math.Sqrt(2*s*s)
		for yIdx := -1; float64(yIdx)*s < height; yIdx++ {
			y := float64(yIdx) * (s + s*math.Sin(gg.Radians(45)))
			dx := x
			if yIdx%2 == 0 {
				dx = x + s*math.Cos(gg.Radians(45))
			}
			drawCube(dc, dx, y, s)
			dc.Stroke()
		}
	}

	// dc.Translate(-minX, -minY)

	dc.SavePNG(namer.NameWithFileType("png"))
}
