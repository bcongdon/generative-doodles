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
	nodes := []rtreego.Spatial{}

	for _, path := range w.paths {
		nodes = append(nodes, path.Nodes[0])
	}

	w.tree = rtreego.NewTree(2, 25, 50, nodes...)
}
