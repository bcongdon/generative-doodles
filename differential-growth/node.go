package main

import (
	"math"

	"github.com/dhconnelly/rtreego"
	"github.com/fogleman/gg"
)

type Node struct {
	X       float64
	Y       float64
	nextX   float64
	nextY   float64
	IsFixed bool
}

func (n *Node) Iterate() {
	if n.IsFixed {
		return
	}

	n.X = Lerp(n.X, n.nextX, 0.5)
	n.Y = Lerp(n.Y, n.nextY, 0.5)
}

func (n *Node) Draw(dc *gg.Context) {
	dc.DrawPoint(n.X, n.Y, 1)
}

func (n *Node) Dist(other *Node) float64 {
	dx := n.X - other.X
	dy := n.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (n *Node) Bounds() *rtreego.Rect {
	point := rtreego.Point{n.X, n.Y}
	return point.ToRect(0.001)
}
