package main

import (
	"fmt"
	"math"
	"math/rand"

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

func pickPointOnCircle(rand *rand.Rand) Point {
	x := 500 - rand.Float64()*1000
	y := math.Sqrt(circleRadius*circleRadius - x*x)
	if rand.Float64() > 0.5 {
		y = -y
	}

	return Point{x, y}
}

func reflectPoint(slope float64, pos Point) Point {
	d := (pos.x + (pos.y * slope)) / (1.0 + (slope * slope))
	x := 2*d - pos.x
	y := 2*d*slope - pos.y

	return Point{x, y}
}

func pickPointOnCircleWithX(x float64) Point {
	y := math.Sqrt(circleRadius*circleRadius - x*x)
	return Point{x, y}
}

func drawCircle(dc *gg.Context, x1, x2 float64, n int, sx, sy float64) {
	pos1 := pickPointOnCircleWithX(x1)
	pos2 := pickPointOnCircleWithX(x2)

	dc.MoveTo(pos1.x, pos1.y)
	dc.SetLineWidth(1)
	fmt.Println(pos1)
	for i := 0; i < n; i++ {
		fmt.Println(pos1, pos2)
		dc.DrawLine(pos1.x, pos1.y, pos2.x, pos2.y)
		dc.Stroke()
		pos3 := reflectPoint(pos2.y/pos2.x, pos1)
		pos1, pos2 = pos2, pos3
		dc.Scale(sx, sy)
	}
}

func main() {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.Translate(500, 500)

	for i := 0; i < 12; i++ {
		drawCircle(dc, 0, 499.9-.05*float64(i), 44, 1-.0001*float64(i), 1-.0001*float64(i))
		dc.Scale(0.65, 0.65)
		dc.Rotate(gg.Radians(22.5))
	}

	// dc.DrawCircle(0, 0, 500)

	dc.SavePNG("2-19-19.png")
}
