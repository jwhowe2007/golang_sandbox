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

func main() {
    fmt.Println("SUM(1,5) =:", summate(1,5))
    fmt.Println("PROD(1,10) =:", product(1,10))
    fmt.Println("8! =:", factorial(8))
    fmt.Println("8C2 =:", combinations(8,3))
    fmt.Println("8P2 =:", permutations(8,3))
    fmt.Println("H(100) =:", heavisideSimple(100))
    fmt.Println("H(-7) =:", heavisideSimple(-7))
    fmt.Println(powLimited(3, 2, 10), powLimited(3, 3, 20))
    fmt.Println(approxSqrt(1.2))
    fmt.Println(approxCubeRoot(12))
    fmt.Println("1.23779 rounded to the nearest hundredth =:", roundToPlace(1.23779, 2, 10))
}
