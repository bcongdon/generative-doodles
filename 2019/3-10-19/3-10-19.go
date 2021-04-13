package main

import (
	"fmt"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

var namer = fn.New()

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func mapPoint(x1, x2, y1, y2, xt float64) float64 {
	return y1 + (xt-x1)*(y2-y1)/(x2-x1)
}

var drawScale = 10.0

func drawEarlyStopRun(dc *gg.Context, samples []int, cutoff int) {
	bestKnown := samples[0]
	for i := 1; i < cutoff; i++ {
		if samples[i] > bestKnown {
			bestKnown = samples[i]
		}
	}

	chosenIdx := len(samples) - 1
	for i := cutoff; i < len(samples); i++ {
		if samples[i] > bestKnown {
			chosenIdx = i
			break
		}
	}

	maxVal := float64(len(samples))

	dc.MoveTo(0, (maxVal-float64(samples[0]))*drawScale)
	for i, val := range samples {
		dc.LineTo(float64(i)*drawScale, (maxVal-float64(val))*drawScale)
		dc.LineTo(float64(i+1)*drawScale, (maxVal-float64(val))*drawScale)
	}
	dc.Stroke()

	// Chosen line (vertical)
	dc.DrawLine((float64(chosenIdx)+0.5)*drawScale, 0, (float64(chosenIdx)+0.5)*drawScale, maxVal*drawScale)

	// Best known line (horizontal)
	dc.DrawLine(0, (maxVal-float64(bestKnown))*drawScale, float64(len(samples))*drawScale, (maxVal-float64(bestKnown))*drawScale)
	fmt.Println(samples, bestKnown, chosenIdx, samples[chosenIdx])
}

func main() {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	// dc := gg.NewContext(width, height)
	dc := gg.NewContext(80*12, 80*12)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(1)
	// dc.Translate(width/2, height/2)
	dc.Translate((80*2)/2, 15)

	input := []int{1, 2, 3, 4, 5}

	for idx, ordering := range permutations(input) {
		x := idx / 12
		y := idx % 12
		dc.Push()
		dc.Translate(float64(x)*float64(len(input)+3)*drawScale, float64(y)*float64(len(input)+3)*drawScale)
		drawEarlyStopRun(dc, ordering, 2)
		dc.Pop()
	}

	dc.Stroke()
	// dc.Translate(-minX, -minY)

	dc.SavePNG(namer.NameWithFileType("png"))
}
