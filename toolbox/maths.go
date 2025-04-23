package toolbox

import (
	"fmt"
	"math"
)

func Summate(init, bound int) int {
	sum := 0

	// For loop - note that init and step/post are both optional.
	for i := init; i <= bound; i++ {
		sum += i
	}

	return sum
}

func Product(init, bound int) int {
	product := 1

	for i := init; i < bound; i++ {
		product *= i
	}

	return product
}

func Factorial(bound int) int {
	factorial := 1

	for i := 1; i <= bound; i++ {
		factorial *= i
	}

	return factorial
}

func Permutations(count, group int) int {
	return (Factorial(count)) / (Factorial(count - group))
}

func Combinations(count, group int) int {
	return Permutations(count, group) / Factorial(group)
}

func HeavisideSimple(x int) int {
	if x >= 0 {
		return 1
	} else {
		return 0
	}
}

func PowLimited(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func RoundToPlace(num, place, base float64) float64 {
	basePosition := math.Pow(base, place)

	if num == 0 {
		return 0
	}

	return math.Floor((num+1/(2*basePosition))*basePosition) / basePosition
}

func ApproxSqrt(x float64) float64 {
	var z float64 = x / 2.0
	var epsilon float64 = 0.0000000001 // Result should be within 1 billionth of actual

	// Basic limiting function - test func must be within ε of the given value
	for math.Abs(z*z-x) > epsilon {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func ApproxCubeRoot(x float64) float64 {
	var z float64 = x / 3.0
	var epsilon float64 = 0.0000000001

	for math.Abs(z*z*z-x) > epsilon {
		z -= (z*z*z - x) / (3 * z * z)
	}

	return z
}

// NB: this function is a generalization of the algorithm defined in the above two root-seeking functions and doesn't account for multiple roots, weird corner cases, etc. This works best for algebraic functions with real roots, for example.
func NaiveNthRoot(x, n float64) float64 {
	var z float64 = x / n
	var epsilon float64 = 0.0000000001

	for math.Abs(math.Pow(z, n)-x) > epsilon {
		z -= (math.Pow(z, n) - x) / (n * math.Pow(z, n-1))
	}

	return z
}

func CosCubeRoot() float64 {
	var z float64 = 1
	var epsilon float64 = 0.00000000001

	for math.Abs((math.Cos(z) - math.Pow(z, 3))) > epsilon {
		z -= (math.Cos(z) - math.Pow(z, 3)) / (-math.Sin(z) - 3*z*z)
	}

	return z
}

func WeirdRoot() float64 {
	var z float64 = .5
	var epsilon float64 = 0.0000000001
	i := 0

	for math.Abs(z/ApproxSqrt(1+z*z)) > epsilon {
		i++
		z -= (z / ApproxSqrt(1+z*z)) / ((2*z*z - z + 2) / 2 * (math.Pow(ApproxSqrt(1+z*z), 3)))
	}

	// Infinitessimally small roots (less than epsilon) should be truncated to zero.
	if z < epsilon {
		z = 0
	}

	fmt.Println("iterations :=", i)

	return z
}

type DerivativeFunction func(x float64) float64

func Derivative(x float64, dfunc DerivativeFunction) float64 {
	var epsilon float64 = 0.0000001
	var h float64 = 1
	var differential float64

	for h > epsilon {
		differential = (dfunc(x+h) - dfunc(x)) / h
		h /= 10
	}

	return RoundToPlace(differential, 4, 10)
}
