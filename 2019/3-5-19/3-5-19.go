package main

import (
	"math"
	"math/rand"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
	"github.com/ojrac/opensimplex-go"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

var namer = fn.New()

var noise = opensimplex.New(0)

func mapPoint(x1, x2, y1, y2, xt float64) float64 {
	return y1 + (xt-x1)*(y2-y1)/(x2-x1)
}

func drawCircleRing(dc *gg.Context, r1, r2 float64) {
	circ := math.Pi * 2 * (r1 + r2) / 2
	segments := int(math.Floor((circ / 25) + rand.Float64()*10))

	aOffset := rand.Float64() * math.Pi
	segmentArc := 2 * math.Pi / float64(segments)
	for aIdx := 0; aIdx < segments; aIdx++ {
		dc.Push()
		a := segmentArc*float64(aIdx) + aOffset
		dc.MoveTo(r1*math.Cos(a+aOffset),
			r1*math.Sin(a+aOffset))
		dc.QuadraticTo(
			(r1+(r2-r1)/2)*math.Cos(a+aOffset)+(r1*0.25)*noise.Eval2(r2, a),
			(r1+(r2-r1)/2)*math.Sin(a+aOffset)+(r1*0.10)*noise.Eval2(r1, a),
			r2*math.Cos(a+aOffset),
			r2*math.Sin(a+aOffset))
		dc.Stroke()

		// dr2 := r2 + 20*noise.Eval2(r2, a)
		// effectiveA2 := a + segmentArc
		// if aIdx == segments-1 {
		// 	effectiveA2 = aOffset
		// }

		// dc.Shear(0.2*noise.Eval2(r1, r2), 0.2*noise.Eval2(r1, r2))

		// dr3 := r2 + 20*noise.Eval2(r2, effectiveA2)
		// dc.DrawLine(
		// 	dr2*math.Cos(a),
		// 	dr2*math.Sin(a),
		// 	dr3*math.Cos(a+segmentArc),
		// 	dr3*math.Sin(a+segmentArc))
		dc.Pop()
	}
	dc.Stroke()
	dc.SetLineWidth(0.5)
	// dc.DrawCircle(0, 0, r1)
	// dc.DrawCircle(0, 0, r2)
	dc.Stroke()
	dc.SetLineWidth(1.5)
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

	dc.Translate(300, 300)
	dc.Scale(1.65, 1.65)
	radiusDiff := 5.0
	for r1 := 10.0; r1 < float64(width); r1 += radiusDiff {
		drawCircleRing(dc, r1, r1+radiusDiff)
		radiusDiff *= 1.05
	}

	dc.Stroke()
	// dc.Translate(-minX, -minY)

	dc.SavePNG(namer.NameWithFileType("png"))
}
