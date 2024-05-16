package main

import (
    "fmt"
    "math"
)

func summate(init, bound int) int {
    sum := 0

    // For loop - note that init and step/post are both optional.
    for i := init; i <= bound; i++ {
        sum += i
    }

    return sum
}

func product(init, bound int) int {
    product := 1

    for i := init; i < bound; i++ {
        product *= i
    }

    return product
}

func factorial(bound int) int {
    factorial := 1

    for i := 1; i <= bound; i++ {
        factorial *= i
    }

    return factorial
}

func permutations(count, group int) int {
    return (factorial(count)) / (factorial(count-group))
}

func combinations(count, group int) int {
    return permutations(count, group) / factorial(group)
}

func heavisideSimple(x int) int {
    if x >= 0 {
        return 1
    } else {
        return 0
    }
}

func powLimited(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func roundToPlace(num, place, base float64) float64 {
    basePosition := math.Pow(base, place)

    if num == 0	{
        return 0
    }

    return math.Floor((num + 1/(2*basePosition)) * basePosition) / basePosition
}

func approxSqrt(x float64) float64 {
    var z float64 = x/2.0
    var epsilon float64 = 0.0000000001 // Result should be within 1 billionth of actual

    // Basic limiting function - test func must be within ε of the given value
    for math.Abs(z*z-x) > epsilon {
        z -= (z*z-x) / (2*z)
    }

    return z
}

func approxCubeRoot(x float64) float64 {
    var z float64 = x/3.0
    var epsilon float64 = 0.0000000001

    for math.Abs(z*z*z-x) > epsilon {
        z -= (z*z*z-x) / (3*z*z)
    }

    return z
}

// NB: this function is a generalization of the algorithm defined in the above two root-seeking functions and doesn't account for multiple roots, weird corner cases, etc. This works best for algebraic functions with real roots, for example.
func naiveNthRoot(x, n float64) float64 {
    var z float64 = x / n
    var epsilon float64 = 0.0000000001

    for math.Abs(math.Pow(z,n) - x) > epsilon {
        z -= (math.Pow(z,n) - x) / (n * math.Pow(z,n-1))
    }

    return z
}

func cosCubeRoot() float64 {
    var z float64 = 1
    var epsilon float64 = 0.00000000001

    for math.Abs((math.Cos(z)-math.Pow(z,3))) > epsilon {
        z -= ((math.Cos(z)-math.Pow(z,3))) / (-math.Sin(z) - 3*z*z)
    }

    return z
}

func weirdRoot() float64 {
    var z float64 = .5
    var epsilon float64 = 0.0000000001
    i := 0

    for math.Abs(z/approxSqrt(1+z*z)) > epsilon {
        i++
        z -= (z/approxSqrt(1+z*z)) / ((2*z*z-z+2)/2*(math.Pow(approxSqrt(1+z*z), 3)))
    }

    // Infinitessimally small roots (less than epsilon) should be truncated to zero.
    if z < epsilon {
        z = 0
    }

    fmt.Println("iterations :=", i)

    return z
}

type DerivativeFunction func(x float64) float64

func derivative(x float64, dfunc DerivativeFunction) float64 {
    var epsilon float64 = 0.0000001
    var h float64 = 1
    var differential float64

    for h > epsilon {
        differential = (dfunc(x+h) - dfunc(x)) / h
        h /= 10
    }

    return roundToPlace(differential, 4, 10)
}

func main() {
    fmt.Println("SUM(1,5) =:", summate(1,5))
    fmt.Println("PROD(1,10) =:", product(1,10))
    fmt.Println("8! =:", factorial(8))
    fmt.Println("8C2 =:", combinations(8,3))
    fmt.Println("8P2 =:", permutations(8,3))
    fmt.Println("H(100) =:", heavisideSimple(100))
    fmt.Println("H(-7) =:", heavisideSimple(-7))
    fmt.Println(powLimited(3, 2, 10), powLimited(3, 3, 20))
    fmt.Println("Square root of 1.2 :=", approxSqrt(1.2))
    fmt.Println("Cube root of 12 :=", approxCubeRoot(12))
    fmt.Println("7th root of 23 :=", naiveNthRoot(23,7))
    fmt.Println("5th root of 0.5 :=", naiveNthRoot(0.5,5))
    fmt.Println("Cosine cube root (x=cos^3(x)) :=", cosCubeRoot())
    fmt.Println("Weird root :=", weirdRoot())
    fmt.Println("1.23779 rounded to the nearest hundredth :=", roundToPlace(1.23779, 2, 10))

    var derivFunc DerivativeFunction = func (x float64) float64 {
        return math.Hypercos(x) // Compute derivative of log_10(x) using natural logs (1/x*ln(10))
    }

    fmt.Println("Derivative of log_10(x) at x = 100:", derivative(100, derivFunc))
}
