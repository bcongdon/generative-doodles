package main

import (
	"math"

	"github.com/dhconnelly/rtreego"
)

func Lerp(start, stop, amt float64) float64 {
	return amt*(stop-start) + start
}

func MakeRadiusFilter(source *Node, radius float64) rtreego.Filter {
	return func(results []rtreego.Spatial, object rtreego.Spatial) (refuse, abort bool) {
		other := object.(*Node)
		if source.Dist(other) > radius {
			refuse = true
		} else {
			refuse = false
		}
		return
	}
}

func ClampAbs(x, clamp float64) float64 {
	if math.Abs(x) > clamp {
		return x / math.Abs(x) * clamp
	}

	return x
}
