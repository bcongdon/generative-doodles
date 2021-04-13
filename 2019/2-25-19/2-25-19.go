package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

var namer = fn.New()

func genGosperSequence(levels int) string {
	replace := func(c rune) string {
		switch c {
		case 'a':
			return "a-b--b+a++aa+b-"
		case 'b':
			return "+a-bb--b-a++a+b"
		default:
			return string(c)
		}
	}

	output := "a"
	for level := 0; level < levels; level++ {
		newOutput := make([]string, len(output))

		for idx, c := range output {
			newOutput[idx] = replace(c)
		}

		output = strings.Join(newOutput, "")
	}
	return output
}

func iterateGosperSequence(seq string, f func(x, y float64, idx int)) {
	x, y, a := 0.0, 0.0, 0.0
	for idx, c := range seq {
		switch c {
		case 'a', 'b':
			f(x, y, idx)
			x += 7 * math.Cos(gg.Radians(a))
			y += 7 * math.Sin(gg.Radians(a))
		case '-':
			a += 60
		case '+':
			a -= 60
		}
	}
}

func mapPoint(x1, x2, y1, y2, xt float64) float64 {
	return y1 + (xt-x1)*(y2-y1)/(x2-x1)
}

func main() {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(1)

	sequence := genGosperSequence(6)

	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64
	iterateGosperSequence(sequence, func(x, y float64, _ int) {
		minX = math.Min(x, minX)
		minY = math.Min(y, minY)
		maxX = math.Max(x, maxX)
		maxY = math.Max(y, maxY)
	})

	fmt.Println(maxX, maxY, minX, minY)

	px, py := 0.0, 0.0
	iterateGosperSequence(sequence, func(x, y float64, idx int) {
		nx := mapPoint(minX, maxX, 0, width, x)
		ny := mapPoint(minY, maxY, 0, height, y)
		if idx%2 == 0 || idx%3 == 0 || idx%5 == 0 || idx%7 == 0 || idx%11 == 0 || idx%13 == 0 {
			px, py = nx, ny
			return
		}

		// dc.SetLineWidth(1.0 + math.Max(0, (1./5e5)*(nx-500)*(nx-500)+(1./5e5)*(ny-500)*(ny-500)))
		dc.SetLineWidth(1.75 + 0.5*math.Cos(6*math.Pi*2*float64(idx)/float64(len(sequence))))
		dc.QuadraticTo(px, py, nx, ny)
		// dc.LineTo(nx, ny)
		dc.Stroke()
		dc.MoveTo(nx, ny)
		px, py = nx, ny
	})
	dc.Stroke()
	// dc.Translate(-minX, -minY)

	dc.SavePNG(namer.NameWithFileType("png"))
}
