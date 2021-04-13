package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const width = 1000
const height = 1000
const circleRadius = 500.0

type Point struct {
	x float64
	y float64
}

func lineCircleIntersection(p1, p2 Point) (Point, Point) {
	D := p1.x*p2.y - p2.x*p1.y
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	dr := math.Sqrt(dx*dx + dy*dy)

	sgnDy := 1.0
	if dy < 0 {
		sgnDy = -1.0
	}

	xIntersect1 := (D*dy + sgnDy*dx*math.Sqrt(circleRadius*circleRadius*dr*dr-(D*D))) / (dr * dr)
	xIntersect2 := (D*dy - sgnDy*dx*math.Sqrt(circleRadius*circleRadius*dr*dr-(D*D))) / (dr * dr)

	yIntersect1 := ((-D * dx) + math.Abs(dy)*math.Sqrt(circleRadius*circleRadius*dr*dr-(D*D))) / (dr * dr)
	yIntersect2 := ((-D * dx) - math.Abs(dy)*math.Sqrt(circleRadius*circleRadius*dr*dr-(D*D))) / (dr * dr)

	return Point{xIntersect1, yIntersect1}, Point{xIntersect2, yIntersect2}
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.Translate(500, 500)

	fib := []float64{1, 1, 2, 3, 5, 8, 13}
	for i, elem := range fib {
		fib[i] = elem * 30
	}
	dc.SetDash(fib...)

	pos := Point{x: 0, y: 0}
	for i := 0; i < 750; i++ {
		angle := r1.Float64() * 180
		slope := math.Tan(gg.Radians(angle))
		pos2 := Point{x: pos.x + 100, y: pos.y + 100*slope}

		intersect1, intersect2 := lineCircleIntersection(pos, pos2)
		fmt.Println(intersect1, intersect2)

		dc.SetLineWidth(1)
		dc.DrawLine(intersect1.x, intersect1.y, intersect2.x, intersect2.y)
		dc.Stroke()

		if !math.IsNaN(intersect2.x) {
			pos = pos2
		} else {
			nx := 250 - r1.Float64()*500
			ny := 250 - r1.Float64()*500
			pos = Point{x: nx, y: ny}
		}
	}

	dc.SavePNG("2-18-19.png")
}
