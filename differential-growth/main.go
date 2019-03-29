package main

import (
	"fmt"
	"os"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

var namer = fn.New()

func main() {
	width, height := 1000.0, 1000.0
	world := NewWorld(width, height)

	nodes := make([]*Node, 50)
	for i := 0; i < 50; i++ {
		nodes[i] = NewNode(500, float64(i)*20)
	}
	path := NewPath(nodes)
	path.UseBrownianMotion = true
	path.NodeInjectionInterval = 10

	world.paths = append(world.paths, path)

	dirName := namer.Name()
	os.Mkdir(dirName, os.ModePerm)

	for i := 0; i < 1000; i++ {
		fmt.Println(i)

		dc := gg.NewContext(int(width), int(height))
		dc.SetRGB(245./255, 245./255, 245./255)
		dc.Clear()
		dc.SetRGB(0, 0, 0)

		world.Draw(dc)
		dc.Stroke()

		dc.SavePNG(fmt.Sprintf("%s/%d.png", dirName, i))

		world.Iterate()
	}
}
