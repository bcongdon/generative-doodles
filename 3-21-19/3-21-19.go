package main

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

var namer = fn.New()

type Car int

const (
	Empty = iota
	Right
	Down
)

func randomDirection() Car {
	if rand.Float64() < .5 {
		return Right
	}
	return Down
}

func simulateTraffic(density float64, iterations int, width, height int) [][]Car {
	grid := make([][]Car, width)
	for x := 0; x < width; x++ {
		grid[x] = make([]Car, height)

		for y := 0; y < height; y++ {
			if rand.Float64() < density {
				grid[x][y] = randomDirection()
			}
		}
	}

	palette := make([]color.Color, 3)
	palette[0] = color.Gray{255 - 0}
	palette[1] = color.Gray{255 - 50}
	palette[2] = color.Gray{255 - 125}
	rect := image.Rect(0, 0, width, height)

	frames := []*image.Paletted{}

	for i := 0; i < iterations; i++ {
		img := image.NewPaletted(rect, palette)

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				switch grid[x][y] {
				case Right:
					nx := (x - 1 + len(grid)) % len(grid)
					if grid[nx][y] == Empty {
						grid[nx][y] = grid[x][y]
						grid[x][y] = Empty
					}
				case Down:
					ny := (y - 1 + len(grid[x])) % len(grid[x])
					if grid[x][ny] == Empty {
						grid[x][ny] = grid[x][y]
						grid[x][y] = Empty
					}
				}
			}
		}

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				img.SetColorIndex(width-x-1, height-y-1, uint8(grid[x][y]))
			}
		}
		frames = append(frames, img)
	}

	delays := make([]int, len(frames))
	for i := 0; i < len(delays); i++ {
		delays[i] = 5
	}

	// anim := gif.GIF{Delay: delays, Image: frames}
	// gif.EncodeAll(os.Stdout, &anim)

	return grid
}

func main() {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(0.5)
	// dc.Translate(width/2, height/2)

	gridWidth := 144
	gridHeight := 89

	grid := simulateTraffic(0.26, 5000, gridWidth, gridHeight)
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] != Empty {
				switch grid[x][y] {
				case Right:
					dc.SetRGB(150./255, 150./255, 150./255)
				case Down:
					dc.SetRGB(50./255, 50./255, 50./255)
				}
				dc.DrawRectangle(
					width/float64(gridWidth)*float64(x),
					height/float64(gridHeight)*float64(y),
					width/float64(gridWidth),
					height/float64(gridHeight))
				dc.Fill()
				dc.Stroke()
			}
		}
	}

	dc.Stroke()

	dc.SavePNG(namer.NameWithFileType("png"))
}
