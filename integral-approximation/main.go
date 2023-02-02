package main

import (
	"fmt"
	"math"
)

var (
	a        = 0
	b        = 1
	n        = 5
	function = func(x float64) float64 {
		return math.Sin(.5 * math.Pow(x, 2))
	}
)

func main() {
	trapazoid(a, b, n, function)
	midpoint(a, b, n, function)
}

func trapazoid(a, b, n int, f func(float64) float64) {
	fmt.Printf("T(%v) = ", n)

	delta := float64((b - a)) / float64(n)
	multiplier := delta / 2
	fmt.Printf("%v(", multiplier)

	var result, v float64
	for i := 0; i < n+1; i++ {
		v = round(f(float64(a) + (float64(i) * delta)))

		if i == 0 {
			result += v
			fmt.Printf("%v + ", v)
		} else if i == n {
			result += v
			fmt.Printf("%v)", v)
		} else {
			result += 2 * v
			fmt.Printf("2(%v) + ", v)
		}
	}

	fmt.Printf(" = %v\n", round(multiplier*result))
}

func midpoint(a, b, n int, f func(float64) float64) {
	fmt.Printf("M(%v) = ", n)

	delta := float64((b - a)) / float64(n)
	fmt.Printf("%v(", delta)

	var result, v float64
	for i := 0; i < n; i++ {
		v = round(f(float64(a) + (float64(i) * delta) + delta/2))
		result += v
		fmt.Printf("%v", v)

		if i == n-1 {
			fmt.Print(")")
		} else {
			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %v\n", round(delta*result))
}

func round(x float64) float64 { return math.Round(x*1000) / 1000 }
