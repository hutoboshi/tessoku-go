package main

import (
	"fmt"
	"math"
)

// RectangularPrism represents a rectangular prism defined by two diagonal points
type RectangularPrism struct {
	x1, y1, z1, x2, y2, z2 float64
}

// NewRectangularPrism creates a new RectangularPrism given two diagonal points
func NewRectangularPrism(a, b, c, d, e, f float64) RectangularPrism {
	return RectangularPrism{
		x1: math.Min(a, d),
		y1: math.Min(b, e),
		z1: math.Min(c, f),
		x2: math.Max(a, d),
		y2: math.Max(b, e),
		z2: math.Max(c, f),
	}
}

// IntersectionVolume calculates the volume of the intersection of two rectangular prisms
func IntersectionVolume(r1, r2 RectangularPrism) float64 {
	x_overlap := math.Max(0, math.Min(r1.x2, r2.x2)-math.Max(r1.x1, r2.x1))
	y_overlap := math.Max(0, math.Min(r1.y2, r2.y2)-math.Max(r1.y1, r2.y1))
	z_overlap := math.Max(0, math.Min(r1.z2, r2.z2)-math.Max(r1.z1, r2.z1))

	return x_overlap * y_overlap * z_overlap
}

// HasPositiveIntersection checks if the intersection volume is positive
func HasPositiveIntersection(r1, r2 RectangularPrism) bool {
	return IntersectionVolume(r1, r2) > 0
}

func main() {
	// Example inputs

	// r1 := NewRectangularPrism(0, 0, 0, 4, 5, 6)
	// r2 := NewRectangularPrism(2, 3, 4, 5, 6, 7)
	r1 := NewRectangularPrism(0, 0, 0, 2, 2, 2)
	r2 := NewRectangularPrism(0, 0, 2, 2, 2, 4)

	if HasPositiveIntersection(r1, r2) {
		fmt.Println("The rectangular prisms have a positive intersection volume.")
	} else {
		fmt.Println("The rectangular prisms do not have a positive intersection volume.")
	}
}
