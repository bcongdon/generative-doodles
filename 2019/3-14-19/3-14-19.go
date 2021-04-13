package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
	"github.com/ojrac/opensimplex-go"
)

var piDigits = "31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491298336733624406566430860213949463952247371907021798609437027705392171762931767523846748184676694051320005681271452635608277857713427577896091736371787214684409012249534301465495853710507922796892589235420199561121290219608640344181598136297747713099605187072113499999983729780499510597317328160963185950244594553469083026425223082533446850352619311881710100031378387528865875332083814206171776691473035982534904287554687311595628638823537875937519577818577805321712268066130019278766111959092164201989"

const width = 1020
const height = 1020
const circleRadius = 500.0

var namer = fn.New()

var noise = opensimplex.New(0)

func lerp(x1, x2, y1, y2, xt float64) float64 {
	return y1 + (xt-x1)*(y2-y1)/(x2-x1)
}

var digitOffsets = [][]int{
	{0, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{1, 1},
}

func main() {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	// dc := gg.NewContext(width, height)
	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(50./255, 50./255, 50./255)
	dc.SetLineWidth(1)
	// dc.Translate(width/2, height/2)

	dc.Stroke()

	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 12); err != nil {
		panic(err)
	}

	padding := 10
	cellSize := 100
	idx := 0
	prevDigit := int64(0)
	for y := 0; y < height-padding*2; y += cellSize {
		for x := 0; x < width-padding*2; x += cellSize {
			digit, _ := strconv.ParseInt(string(piDigits[idx]), 10, 64)

			a1 := lerp(0, 10, 0, 2*math.Pi, float64(prevDigit)) - math.Pi/2
			a2 := lerp(0, 10, 0, 2*math.Pi, float64(digit)) - math.Pi/2
			fmt.Println(x, y, a1, a2)

			if prevDigit == digit {
				dc.DrawCircle(float64(x)+50, float64(y)+50, 25)
			} else {
				dc.DrawArc(float64(x)+50, float64(y)+50, 25, a1, a2)
			}
			dc.DrawPoint(
				float64(x)+50+25*math.Cos(a2),
				float64(y)+50+25*math.Sin(a2),
				2)
			// dc.DrawStringAnchored(fmt.Sprintf("%d", digit), float64(x)+50, float64(y)+50, 0.5, 0.5)

			for j := int64(0); j < digit+1; j++ {
				if j == 0 {
					continue
				}
				offsets := digitOffsets[j-1]
				dx, dy := 5*float64(offsets[0]), 5*float64(offsets[1])
				dc.DrawPoint(float64(x)+50+dx, float64(y)+50+dy, 1)
			}
			if digit == 0 {
				dc.DrawPoint(float64(x)+50, float64(y)+50, 5)
			}

			dc.Stroke()

			prevDigit = digit
			idx++
		}
	}

	dc.SavePNG(namer.NameWithFileType("png"))
}
