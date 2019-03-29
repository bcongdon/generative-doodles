package main

import (
	"math/rand"

	"github.com/dhconnelly/rtreego"
)

type InjectionMode int

const (
	RandomInjection InjectionMode = iota
	CurvatureInjection
)

const minAttractionDistance = 10.0
const attractionForce = 0.001
const repulsionForce = 0.8
const repulsionRadius = 20.0
const alignmentForce = 0.001
const maxEdgeDistance = 30.0
const minEdgeDistance = 20.0

type Path struct {
	Nodes                    []*Node
	UseBrownianMotion        bool
	IsClosed                 bool
	NodeInjectionInterval    int
	NodeInjectionMode        InjectionMode
	iterationsSinceInjection int
}

func NewPath(nodes []*Node) *Path {
	return &Path{
		Nodes: nodes,
	}
}

func (p *Path) Iterate(tree *rtreego.Rtree) {
	for _, node := range p.Nodes {
		if p.UseBrownianMotion {
			p.applyBrownianMotion(node)
		}

		p.applyAttraction(node)
		p.applyRepulsion(node, tree)
		p.applyAlignment(node)

		node.Iterate()
	}

	p.splitEdges()
	p.pruneNodes()

	if p.iterationsSinceInjection >= p.NodeInjectionInterval {
		p.injectNode()
		p.iterationsSinceInjection = 0
	}
	p.iterationsSinceInjection++
}

func (p *Path) applyBrownianMotion(node *Node) {
	node.nextX += -0.5 + rand.Float64()
	node.nextY += -0.5 + rand.Float64()
}

func (p *Path) applyAttraction(node *Node) {
	if node.IsFixed {
		return
	}

	prevNode, nextNode := p.connectedNodes(node)
	for _, connectedNode := range []*Node{prevNode, nextNode} {
		if connectedNode == nil {
			continue
		}
		dist := node.Dist(connectedNode)
		if dist >= minAttractionDistance {
			node.nextX = Lerp(node.nextX, connectedNode.nextX, attractionForce)
			node.nextY = Lerp(node.nextY, connectedNode.nextY, attractionForce)
		}
	}
}

func (p *Path) applyRepulsion(node *Node, tree *rtreego.Rtree) {
	neighbors := tree.NearestNeighbors(100, rtreego.Point{node.X, node.Y}, MakeRadiusFilter(node, repulsionRadius))

	for _, treeNode := range neighbors {
		neighborNode := treeNode.(*Node)
		node.nextX = Lerp(node.nextX, neighborNode.nextX, -repulsionForce)
		node.nextY = Lerp(node.nextY, neighborNode.nextY, -repulsionForce)
	}
}

// TODO
func (p *Path) applyAlignment(node *Node) {
	prevNode, nextNode := p.connectedNodes(node)
	if prevNode != nil && nextNode != nil && !node.IsFixed {
		midpointX, midpointY := prevNode.MidpointTo(nextNode)

		node.nextX = Lerp(node.nextX, midpointX, alignmentForce)
		node.nextY = Lerp(node.nextY, midpointY, alignmentForce)
	}
}

func (p *Path) splitEdges() {
	newNodes := make([]*Node, 0, len(p.Nodes))
	for _, node := range p.Nodes {
		prevNode, _ := p.connectedNodes(node)
		if prevNode != nil && prevNode.Dist(node) >= maxEdgeDistance {
			midpointX, midpointY := node.MidpointTo(prevNode)
			midpointNode := NewNode(midpointX, midpointY)
			newNodes = append(newNodes, midpointNode)
		}
		newNodes = append(newNodes, node)
	}
	p.Nodes = newNodes
}

func (p *Path) pruneNodes() {
	for idx := 0; idx < len(p.Nodes); idx++ {
		node := p.Nodes[idx]
		prevNode, _ := p.connectedNodes(node)
		if prevNode != nil && node.Dist(prevNode) < minEdgeDistance {
			p.Nodes = append(p.Nodes[:idx], p.Nodes[idx+1:]...)
			idx--
		}
	}
}

func (p *Path) injectNode() {
	switch p.NodeInjectionMode {
	case RandomInjection:
		p.injectRandomNode()
	case CurvatureInjection:
		p.injectByCurvature()
	default:
		panic("Unknown injection mode!")
	}
}

func (p *Path) injectRandomNode() {
	index := 1 + rand.Intn(len(p.Nodes)-1)
	injectionNode := p.Nodes[index]
	prevNode, nextNode := p.connectedNodes(injectionNode)

	if prevNode != nil && nextNode != nil && injectionNode.Dist(prevNode) > minEdgeDistance {
		midpointX, midpointY := injectionNode.MidpointTo(prevNode)
		midpointNode := NewNode(midpointX, midpointY)
		p.Nodes = append(p.Nodes[:index], append([]*Node{midpointNode}, p.Nodes[index:]...)...)
	}
}

func (p *Path) injectByCurvature() {
	panic("Unimplemented injection mode")
}

func (p *Path) connectedNodes(node *Node) (prevNode, nextNode *Node) {
	if len(p.Nodes) <= 1 {
		return
	}

	nodeIdx := -1
	for idx, n := range p.Nodes {
		if n == node {
			nodeIdx = idx
			break
		}
	}
	if nodeIdx == -1 {
		return
	}

	if nodeIdx == 0 && p.IsClosed {
		prevNode = p.Nodes[len(p.Nodes)-1]
	} else if nodeIdx > 0 {
		prevNode = p.Nodes[nodeIdx-1]
	}

	if nodeIdx == len(p.Nodes)-1 && p.IsClosed {
		nextNode = p.Nodes[0]
	} else if nodeIdx < len(p.Nodes)-1 {
		nextNode = p.Nodes[nodeIdx+1]
	}
	return
}
