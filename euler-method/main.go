package main

import (
	"fmt"
	"math"
)

// Put the function here:
func F(x float64, y float64) float64 {
	// (dy)/(dx) = 6x^2 - 3x^2y
	return 6*math.Pow(x, 2) - 3*math.Pow(x, 2)*y
}

const (
	// Put the x to evaluate for here:
	x float64 = 1
	// Put the solution for the initial value problem here:
	initial float64 = 3
	// Put the step here:
	h float64 = .001
)

func main() {
	var (
		currentX float64 = 0
		currentY         = initial
	)

	for i := 0; i < int(x/h); i++ {
		currentY = euler(F, currentX, currentY, h)
		currentX += h
	}

	fmt.Println(currentY)
}

func euler(f func(float64, float64) float64, prevx, prevy, h float64) float64 {
	return prevy + h*f(prevx, prevy)
}
