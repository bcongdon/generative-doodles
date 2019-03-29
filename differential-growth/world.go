package main

import (
	"github.com/dhconnelly/rtreego"
	"github.com/fogleman/gg"
)

type World struct {
	tree  *rtreego.Rtree
	paths []*Path
}

func NewWorld() *World {
	world := &World{
		tree:  rtreego.NewTree(2, 25, 50),
		paths: []*Path{},
	}
	return world
}

func (w *World) buildTree() {
	nodes := make([]rtreego.Spatial, 0, 1000)

	for _, path := range w.paths {
		for _, node := range path.Nodes {
			nodes = append(nodes, node)
		}
	}

	w.tree = rtreego.NewTree(2, 25, 50, nodes...)
}

func (w *World) Iterate() {
	w.buildTree()
	for _, path := range w.paths {
		path.Iterate(w.tree)
	}
}

func (w *World) Draw(dc *gg.Context) {
	for _, path := range w.paths {
		for nIdx, node := range path.Nodes {
			node.Draw(dc)

			if nIdx > 0 {
				prevNode := path.Nodes[nIdx-1]
				dc.DrawLine(prevNode.X, prevNode.Y, node.X, node.Y)
			}
			dc.Stroke()
		}
	}
}
