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

func NewNode(x, y float64) *Node {
	return &Node{
		X:       x,
		Y:       y,
		nextX:   x,
		nextY:   y,
		IsFixed: false,
	}
}

func (n *Node) Iterate() {
	if n.IsFixed {
		return
	}

	n.X += ClampAbs(n.nextX-n.X, maxVelocity) / 2
	n.Y += ClampAbs(n.nextY-n.Y, maxVelocity) / 2
}

func (n *Node) Draw(dc *gg.Context) {
	dc.DrawPoint(n.X, n.Y, 1)
}

func (n *Node) Dist(other *Node) float64 {
	dx := n.X - other.X
	dy := n.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (n *Node) MidpointTo(other *Node) (x, y float64) {
	x = (n.X + other.X) / 2
	y = (n.Y + other.Y) / 2
	return
}

func (n *Node) Bounds() *rtreego.Rect {
	point := rtreego.Point{n.X, n.Y}
	return point.ToRect(0.001)
}
