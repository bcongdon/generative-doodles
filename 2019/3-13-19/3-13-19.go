package main

import (
	"math"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
	"github.com/ojrac/opensimplex-go"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

var namer = fn.New()

var noise = opensimplex.New(0)

func main() {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	// dc := gg.NewContext(width, height)
	dc := gg.NewContext(1000, 1000)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(50./255, 50./255, 50./255)
	dc.SetLineWidth(1)
	// dc.Translate(width/2, height/2)

	dc.Stroke()
	// dc.Translate(-minX, -minY)

	dc.Scale(0.75, 0.75)

	spiralCoords := make([]gg.Point, 0)
	for t := 0.0; t < 4.5*math.Pi; t += math.Pi / 32 {
		r := 20 + 45*t
		x := r * math.Cos(t)
		y := r * math.Sin(t)
		spiralCoords = append(spiralCoords, gg.Point{X: x + 0.4*width, Y: y + 0.4*height})
		// dc.DrawPoint(x+0.4*width, y+0.45*height, 45)
		// dc.Stroke()
	}

	numRows := 75.0
	rowHeight := height/numRows - 10

	for y := 0.0; y < height; y += height / numRows {
		startX := 0.0
		for startX < width {
			xWidth := 5 + math.Abs(noise.Eval2(startX, y)*20)

			startPoint := gg.Point{X: startX, Y: y}
			nearSpiral := false
			for _, point := range spiralCoords {
				if startPoint.Distance(point) < 45 {
					nearSpiral = true
				}
			}

			if !nearSpiral {
				dc.MoveTo(startX+2*noise.Eval2(startX, y), y+2*noise.Eval2(y, startX))
				dc.LineTo(startX+xWidth+2*noise.Eval2(startX+xWidth, y), y+2*noise.Eval2(y, startX))
				dc.LineTo(startX+xWidth+2*noise.Eval2(startX+xWidth, y), y+rowHeight+2*noise.Eval2(y, startX))
				dc.LineTo(startX+2*noise.Eval2(startX, y), y+rowHeight+2*noise.Eval2(y, startX))

				// dc.DrawRectangle(startX, y, xWidth, rowHeight)
				dc.FillPreserve()
				dc.Stroke()
			}

			startX += xWidth + 5
		}
	}

	dc.SavePNG(namer.NameWithFileType("png"))
}
