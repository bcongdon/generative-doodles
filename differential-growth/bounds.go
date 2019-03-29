package main

type Bounds struct {
	minX float64
	minY float64
	maxX float64
	maxY float64
}

func NewBounds(x, y, w, h float64) Bounds {
	return Bounds{
		minX: x,
		minY: y,
		maxX: x + w,
		maxY: y + h,
	}
}

func (b Bounds) Contains(x, y float64) bool {
	return x >= b.minX && x <= b.maxX && y >= b.minY && y <= b.maxY
}
