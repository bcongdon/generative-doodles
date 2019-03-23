package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

var namer = fn.New()

type Agent struct {
	x               float64
	y               float64
	a               float64
	targetA         float64
	reactivity      int
	reactionTimer   int
	target          *Agent
	positions       []gg.Point
	targetPositions []gg.Point
}

func (a Agent) DistSq(other *Agent) float64 {
	dx := (a.x - other.x)
	dy := (a.y - other.y)
	return dx*dx + dy*dy
}

func (a Agent) AngleTo(other *Agent) float64 {
	dx := (other.x - a.x)
	dy := (other.y - a.y)

	return math.Atan2(dy, dx)
}

func (a *Agent) Move() {
	a.a += 0.05 * (a.targetA - a.a)
	a.x += math.Cos(a.a)
	a.y += math.Sin(a.a)
	a.reactionTimer--
	a.positions = append(a.positions, gg.Point{X: a.x, Y: a.y})
}

func (a *Agent) UpdateTargetPosition() {
	a.targetPositions = append(a.targetPositions, gg.Point{X: a.target.x, Y: a.target.y})
}

func simulateOoda(agents []*Agent, iterations int) [][]gg.Point {
	for iter := 0; iter < iterations; iter++ {
		for i, agent := range agents {
			if agent.reactionTimer <= 0 {
				agent.reactionTimer = agent.reactivity
				closestIdx := -1
				for j, other := range agents {
					if i == j {
						continue
					}
					if closestIdx < 0 || agent.DistSq(other) < agent.DistSq(agents[closestIdx]) {
						closestIdx = j
					}
				}
				agent.target = agents[closestIdx]
				agent.targetA = agent.AngleTo(agents[closestIdx])
			}
			agent.Move()
		}
		for _, agent := range agents {
			agent.UpdateTargetPosition()
		}
	}

	allPositions := make([][]gg.Point, len(agents))
	for i, agent := range agents {
		allPositions[i] = agent.positions
	}

	return allPositions
}

func NewAgent(x, y float64, reactivity int) *Agent {
	return &Agent{
		x:          x,
		y:          y,
		a:          rand.Float64() * math.Pi * 2,
		reactivity: reactivity,
		positions:  []gg.Point{},
	}
}

func mapPoint(x1, x2, y1, y2, xt float64) float64 {
	return y1 + (xt-x1)*(y2-y1)/(x2-x1)
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(0.5)
	// dc.Translate(width/2, height/2)

	agents := []*Agent{}
	for i := 0; i < 5; i++ {
		agents = append(agents, NewAgent(r1.Float64()*100, r1.Float64()*100, 10+int(r1.Float64()*10)))
	}

	paths := simulateOoda(agents, 100)
	minX, maxX := math.MaxFloat64, -math.MaxFloat64
	minY, maxY := math.MaxFloat64, -math.MaxFloat64
	for _, path := range paths {
		for _, point := range path {
			minX = math.Min(minX, point.X)
			maxX = math.Max(maxX, point.X)
			minY = math.Min(minY, point.Y)
			maxY = math.Max(maxY, point.Y)
		}
	}

	for agentIdx, agentPath := range paths {
		cVal := 1 - mapPoint(0, float64(len(paths)), 0.3, 0.9, float64(agentIdx))
		for idx, point := range agentPath {
			dc.SetRGB(cVal, cVal, cVal)
			dc.SetLineWidth(2 + float64(agentIdx)/float64(len(paths)))
			screenX := mapPoint(minX, maxX, 0, width, point.X)
			screenY := mapPoint(minY, maxY, 0, height, point.Y)
			if idx == 0 {
				dc.DrawCircle(screenX, screenY, 3)
				dc.Fill()
				dc.MoveTo(screenX, screenY)

			} else {
				dc.LineTo(screenX, screenY)
				dc.Stroke()
			}
			dc.Push()
			targetPoint := agents[agentIdx].targetPositions[idx]
			tSX := mapPoint(minX, maxX, 0, width, targetPoint.X)
			tSY := mapPoint(minY, maxY, 0, height, targetPoint.Y)
			dc.SetLineWidth(0.5)
			dc.SetRGB(cVal*0.9, cVal*0.9, cVal*0.9)
			dc.DrawLine(screenX, screenY, tSX, tSY)
			dc.Stroke()
			dc.Pop()
			dc.MoveTo(screenX, screenY)
			// dc.DrawPoint(screenX, screenY, 1)
		}
		dc.Stroke()
	}

	dc.Stroke()

	dc.SavePNG(namer.NameWithFileType("png"))
}
