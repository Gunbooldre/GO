package main

import (
	"fmt"
)

// GenDisplaceFn returns a function that calculates displacement based on time
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	// Test 1
	a1, vo1, so1 := 10.0, 2.0, 1.0
	displacementFn1 := GenDisplaceFn(a1, vo1, so1)
	t1 := 3.0
	result1 := displacementFn1(t1)
	fmt.Printf("Test 1: Acceleration=%.2f, Initial Velocity=%.2f, Initial Displacement=%.2f, Time=%.2f\n", a1, vo1, so1, t1)
	fmt.Printf("Resulting Displacement: %.2f\n", result1)

	// Test 2
	a2, vo2, so2 := 5.0, 3.0, 2.0
	displacementFn2 := GenDisplaceFn(a2, vo2, so2)
	t2 := 4.0
	result2 := displacementFn2(t2)
	fmt.Printf("Test 2: Acceleration=%.2f, Initial Velocity=%.2f, Initial Displacement=%.2f, Time=%.2f\n", a2, vo2, so2, t2)
	fmt.Printf("Resulting Displacement: %.2f\n", result2)
}
