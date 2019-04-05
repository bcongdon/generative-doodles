package main

import (
	"github.com/fogleman/gg"
	"github.com/kyroy/kdtree"
)

type World struct {
	tree   *kdtree.KDTree
	paths  []*Path
	bounds Bounds
}

func NewWorld(width, height float64) *World {
	world := &World{
		paths:  []*Path{},
		bounds: NewBounds(0, 0, width, height),
	}
	return world
}

func (w *World) buildTree() {
	nodes := make([]kdtree.Point, 0, 1000)

	for _, path := range w.paths {
		for _, node := range path.Nodes {
			nodes = append(nodes, node)
		}
	}

	// w.tree = rtreego.NewTree(2, 25, 50, nodes...)
	w.tree = kdtree.New(nodes)
}

func (w *World) Iterate() {
	w.buildTree()
	for _, path := range w.paths {
		path.Iterate(w.bounds, w.tree)
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
