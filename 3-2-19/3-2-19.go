package main

import (
	"fmt"
	"math"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

var namer = fn.New()

type state struct {
	x float64
	y float64
	a float64
}

type token struct {
	char  rune
	param float64
}

func charToken(char rune) token {
	return token{char, 0.0}
}

func tokenStr(seq string) []token {
	output := make([]token, len(seq))
	for idx, c := range seq {
		output[idx] = charToken(c)
	}
	return output
}

const R = 1.456

func genFractalPathSeq(levels int) []token {
	replace := func(t token) []token {
		switch t.char {
		case 'A':
			return append(
				[]token{token{'F', 1.0}},
				tokenStr("[+A][-A]")...,
			)
		case 'F':
			return []token{token{'F', t.param * R}}
		default:
			return []token{t}
		}
	}

	output := []token{token{'A', 0.0}}
	for level := 0; level < levels; level++ {
		newOutput := make([]token, 0)

		for _, c := range output {
			newOutput = append(newOutput, replace(c)...)
		}

		output = newOutput
	}
	return output
}

func iterateSequence(seq []token, f func(x, y float64, jump bool, stackLen, idx int)) {
	stack := make([]state, 0)
	x, y, a := 0.0, 0.0, 0.0
	for idx, t := range seq {
		switch t.char {
		case 'F':
			x += 5 * t.param * math.Cos(gg.Radians(a))
			y += 5 * t.param * math.Sin(gg.Radians(a))
			f(x, y, false, len(stack), idx)
		case '-':
			a += 75
		case '+':
			a -= 22.5
		case '[':
			stack = append(stack, state{x, y, a})
		case ']':
			state := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			x = state.x
			y = state.y
			a = state.a
			f(x, y, true, len(stack), idx)
		}
	}
}

func (t token) String() string {
	if t.param < 1e-10 {
		return fmt.Sprintf("%c", t.char)
	}
	return fmt.Sprintf("%c(%0.2f)", t.char, t.param)
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

	sequence := genFractalPathSeq(18)

	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64
	iterateSequence(sequence, func(x, y float64, _ bool, _, _ int) {
		minX = math.Min(x, minX)
		minY = math.Min(y, minY)
		maxX = math.Max(x, maxX)
		maxY = math.Max(y, maxY)
	})

	fmt.Println(maxX, maxY, minX, minY)

	dc.MoveTo(
		mapPoint(minY, maxY, 0, height, 0),
		mapPoint(minX, maxX, 0, width, 0),
	)
	iterateSequence(sequence, func(x, y float64, jump bool, stackLen, idx int) {
		nx := mapPoint(minX, maxX, 0, width, x)
		ny := mapPoint(minY, maxY, 0, height, y)
		// if idx%2 == 0 || idx%3 == 0 || idx%5 == 0 || idx%7 == 0 || idx%11 == 0 || idx%13 == 0 {
		// 	px, py = nx, ny
		// 	return
		// }

		if jump {
			dc.ClearPath()
			dc.MoveTo(ny, nx)
			return
		}

		// dc.SetLineWidth(2.0 / (0.5*float64(stackLen) + 1))
		dc.SetLineWidth(3.0 / (math.Sqrt(float64(stackLen) + 1)))
		// dc.SetLineWidth(1.0 + math.Max(0, (1./5e5)*(nx-500)*(nx-500)+(1./5e5)*(ny-500)*(ny-500)))
		// dc.SetLineWidth(1.75 + 0.5*math.Cos(6*math.Pi*2*float64(idx)/float64(len(sequence))))
		// dc.QuadraticTo(px, py, nx, ny)
		dc.LineTo(ny, nx)
		dc.Stroke()
		dc.MoveTo(ny, nx)
	})
	dc.Stroke()
	// dc.Translate(-minX, -minY)

	dc.SavePNG(namer.NameWithFileType("png"))
}
