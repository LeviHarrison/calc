// Evaluates a given recursive function (defined in a_n1) at a given
// number (defined in n) using the value of the first number of the
// sequence (defined in a_1). Rounded to the 5th decimal place.
package main

import "fmt"

// a₁
const a_1 = 1

// aₙ₊₁
func a_n1(a_n float64 /* aₙ */) float64 {
	return 1 / (1 + a_n)
}

// n to evaluate to
const n = 10

func main() {
	eval(n)
}

func eval(n int) (r float64) {
	if n == 1 {
		return a_1
	}
	r = a_n1(eval(n - 1))
	fmt.Printf("%.5f\n", r)
	return
}
