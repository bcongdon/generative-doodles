package main

import "github.com/dhconnelly/rtreego"

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
