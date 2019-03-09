package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/bcongdon/fn"
)

var namer = fn.New()

var numColors uint32 = 1

func getClosestColor(greyVal float64) color.Gray16 {
	gradationWidth := math.MaxUint16 / numColors

	return color.Gray16{uint16((uint32(greyVal) / gradationWidth) * gradationWidth)}
}

func make2dFloatSlice(w, h int) [][]float64 {
	a := make([][]float64, w)
	for i := range a {
		a[i] = make([]float64, h)
	}
	return a
}

func main() {
	infile, err := os.Open("building.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer infile.Close()

	img, err := png.Decode(infile)
	if err != nil {
		log.Fatalln(err)
	}

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	newImage := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})

	floatImg := make2dFloatSlice(w, h)
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			floatImg[x][y] += float64(r+g+b) / 3
			newColor := getClosestColor(floatImg[x][y])

			pixelError := floatImg[x][y] - float64(newColor.Y)
			floatImg[x+1][y] += pixelError * 7 / 16
			floatImg[x-1][y+1] += pixelError * 3 / 16
			floatImg[x][y+1] += pixelError * 5 / 16
			floatImg[x+1][y+1] += pixelError * 1 / 16

			newImage.Set(x, y, newColor)
		}
	}

	outfile, err := os.Create(namer.NameWithFileType("png"))
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}
	defer outfile.Close()
	png.Encode(outfile, newImage)
}
