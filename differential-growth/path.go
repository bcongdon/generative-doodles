package main

import (
	"math/rand"

	"github.com/dhconnelly/rtreego"
)

const minAttractionDistance = 10.0
const attractionForce = 0.2
const repulsionForce = 0.2
const repulsionRadius = 100
const alignmentForce = 0.1
const maxEdgeDistance = 100

type Path struct {
	Nodes                    []*Node
	UseBrownianMotion        bool
	IsClosed                 bool
	NodeInjectionInterval    int
	iterationsSinceInjection int
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

	if p.iterationsSinceInjection >= p.NodeInjectionInterval {
		p.injectNode()
	}
	p.iterationsSinceInjection++
}

func (p *Path) applyBrownianMotion(node *Node) {
	node.X += rand.Float64()
	node.Y += rand.Float64()
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
		if prevNode != nil && prevNode.Dist(node) < maxEdgeDistance {
			midpointX, midpointY := node.MidpointTo(prevNode)
			midpointNode := NewNode(midpointX, midpointY)
			newNodes = append(newNodes, midpointNode)
		}
		newNodes = append(newNodes, node)
	}
}

// TODO
func (p *Path) pruneNodes() {

}

// TODO
func (p *Path) injectNode() {

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
	} else if nodeIdx >= 0 {
		prevNode = p.Nodes[nodeIdx-1]
	}

	if nodeIdx == len(p.Nodes)-1 && p.IsClosed {
		nextNode = p.Nodes[0]
	} else if nodeIdx < len(p.Nodes)-1 {
		nextNode = p.Nodes[nodeIdx+1]
	}
	return
}
