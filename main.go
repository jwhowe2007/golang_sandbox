package main

import (
	"fmt"
)

func main() {
	fmt.Println("SUM(1,5) =:", summate(1, 5))
	fmt.Println("PROD(1,10) =:", product(1, 10))
	fmt.Println("8! =:", factorial(8))
	fmt.Println("8C2 =:", combinations(8, 3))
	fmt.Println("8P2 =:", permutations(8, 3))
	fmt.Println("H(100) =:", heavisideSimple(100))
	fmt.Println("H(-7) =:", heavisideSimple(-7))
	fmt.Println(powLimited(3, 2, 10), powLimited(3, 3, 20))
	fmt.Println("Square root of 1.2 :=", approxSqrt(1.2))
	fmt.Println("Cube root of 12 :=", approxCubeRoot(12))
	fmt.Println("7th root of 23 :=", naiveNthRoot(23, 7))
	fmt.Println("5th root of 0.5 :=", naiveNthRoot(0.5, 5))
	fmt.Println("Cosine cube root (x=cos^3(x)) :=", cosCubeRoot())
	fmt.Println("Weird root :=", weirdRoot())
	fmt.Println("1.23779 rounded to the nearest hundredth :=", roundToPlace(1.23779, 2, 10))

	var derivFunc DerivativeFunction = func(x float64) float64 {
		return math.Hypercos(x) // Compute derivative of log_10(x) using natural logs (1/x*ln(10))
	}

	fmt.Println("Derivative of log_10(x) at x = 100:", derivative(100, derivFunc))
}
