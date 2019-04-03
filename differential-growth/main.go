package main

import (
	"fmt"
	"math"
	"os"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

var namer = fn.New()

func main() {
	width, height := 1000.0, 1000.0
	world := NewWorld(width, height)

	numNodes := 25
	nodes := make([]*Node, numNodes)
	for i := 0; i < numNodes; i++ {
		a := float64(i) * ((math.Pi * 2) / float64(numNodes))
		nodes[i] = NewNode(500+100*math.Cos(a), 500+100*math.Sin(a))
	}
	path := NewPath(nodes)
	path.IsClosed = true
	path.UseBrownianMotion = true
	path.NodeInjectionInterval = 1

	world.paths = append(world.paths, path)

	dirName := namer.Name()
	os.Mkdir(dirName, os.ModePerm)

	for i := 0; i < 750; i++ {
		fmt.Println(i, len(path.Nodes))

		dc := gg.NewContext(int(width), int(height))
		dc.SetRGB(245./255, 245./255, 245./255)
		dc.Clear()
		dc.SetRGB(0, 0, 0)

		if i%10 == 0 {
			world.Draw(dc)
			dc.Stroke()
			dc.SavePNG(fmt.Sprintf("%s/%d.png", dirName, i))
		}

		world.Iterate()
		for _, p := range path.Nodes {
			x := p.X - 500
			y := p.Y - 500
			if (x*x)+(y*y) > 1000*1000 {
				p.IsFixed = true
			}
		}
	}
}
