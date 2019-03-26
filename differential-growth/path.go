package main

import "math/rand"

const minAttractionDistance = 10.0
const attractionForce = 0.2

type Path struct {
	Nodes             []*Node
	UseBrownianMotion bool
	IsClosed          bool
}

func (p *Path) Iterate() {
	for _, node := range p.Nodes {
		if p.UseBrownianMotion {
			p.applyBrownianMotion(node)
		}

		p.applyAttraction(node)
	}
}

func (p *Path) applyBrownianMotion(node *Node) {
	node.X += rand.Float64()
	node.Y += rand.Float64()
}

func (p *Path) applyAttraction(node *Node) {
	if node.IsFixed {
		return
	}
	for _, connectedNode := range p.connectedNodes(node) {
		dist := node.Dist(connectedNode)
		if dist >= minAttractionDistance {
			node.nextX = Lerp(node.nextX, connectedNode.nextX, attractionForce)
			node.nextY = Lerp(node.nextY, connectedNode.nextY, attractionForce)
		}
	}
}

func (p *Path) connectedNodes(node *Node) []*Node {
	if len(p.Nodes) <= 1 {
		return []*Node{}
	}

	nodeIdx := -1
	for idx, n := range p.Nodes {
		if n == node {
			nodeIdx = idx
			break
		}
	}
	if nodeIdx == -1 {
		return []*Node{}
	}

	var prevNode, nextNode *Node
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

	connected := make([]*Node, 0)
	if prevNode != nil {
		connected = append(connected, prevNode)
	}
	if nextNode != nil {
		connected = append(connected, nextNode)
	}
	return connected
}
