package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float64)
	RenderSquare(side float64)
}

/* VectorRenderer is a ConcreteImplementor that draws using vector
   primitives (lines and curves) */
type VectorRenderer struct{}

func (v * VectorRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing a circle with radius of %.1f using vector\n", radius)
}

func (v *VectorRenderer) RenderSquare(side float64) {
	fmt.Printf("Drawin a square with side %.1f using vector")
}

// RasterRenderer is a ConcreteImplementor that draws using pixels
type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing a circle with radius of %.1f using raster\n", radius)
}

func (r *RasterRenderer) RenderSquare(side float64) {
	fmt.Printf("Drawing a square with side %.1f using raster\n", side)
}