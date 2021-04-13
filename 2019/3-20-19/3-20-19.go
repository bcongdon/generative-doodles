package main

import (
	"math"

	"github.com/bcongdon/fn"
	"github.com/fogleman/ln/ln"
	"github.com/ojrac/opensimplex-go"
)

var namer = fn.New()
var noise = opensimplex.New(0)

func plotFunc(x, y float64) float64 {
	return 1 - math.Abs(x+y) - math.Abs(y-x)
	return math.Sin(3*x) * math.Cos(3*y) / 10
}

func main() {
	// create a scene and add a single cube
	scene := ln.Scene{}
	box := ln.Box{ln.Vector{-2, -2, -4}, ln.Vector{2, 2, 4}}
	scene.Add(ln.NewFunction(plotFunc, box, ln.Below))

	// define camera parameters
	eye := ln.Vector{3, 0, 3}      // camera position
	center := ln.Vector{1.1, 0, 0} // camera looks at
	up := ln.Vector{0, 0, 1}       // up direction

	// define rendering parameters
	width := 1024.0  // rendered width
	height := 1024.0 // rendered height
	fovy := 50.0     // vertical field of view, degrees
	znear := 0.01    // near z plane
	zfar := 100.0    // far z plane
	step := 0.01     // how finely to chop the paths for visibility testing

	// compute 2D paths that depict the 3D scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	// render the paths in an image
	paths.WriteToPNG(namer.NameWithFileType("png"), width, height)
}
