package main

import (
	"fmt"
)

func main() {
	fmt.Println("SUM(1,5) =:", Summate(1, 5))
	fmt.Println("PROD(1,10) =:", Product(1, 10))
	fmt.Println("8! =:", Factorial(8))
	fmt.Println("8C2 =:", Combinations(8, 3))
	fmt.Println("8P2 =:", Permutations(8, 3))
	fmt.Println("H(100) =:", HeavisideSimple(100))
	fmt.Println("H(-7) =:", HeavisideSimple(-7))
	fmt.Println(PowLimited(3, 2, 10), PowLimited(3, 3, 20))
	fmt.Println("Square root of 1.2 :=", ApproxSqrt(1.2))
	fmt.Println("Cube root of 12 :=", ApproxCubeRoot(12))
	fmt.Println("7th root of 23 :=", NaiveNthRoot(23, 7))
	fmt.Println("5th root of 0.5 :=", NaiveNthRoot(0.5, 5))
	fmt.Println("Cosine cube root (x=cos^3(x)) :=", CosCubeRoot())
	fmt.Println("Weird root :=", WeirdRoot())
	fmt.Println("1.23779 rounded to the nearest hundredth :=", RoundToPlace(1.23779, 2, 10))

	var derivFunc DerivativeFunction = func(x float64) float64 {
		return math.Hypercos(x) // Compute derivative of log_10(x) using natural logs (1/x*ln(10))
	}

	fmt.Println("Derivative of log_10(x) at x = 100:", Derivative(100, derivFunc))
}
