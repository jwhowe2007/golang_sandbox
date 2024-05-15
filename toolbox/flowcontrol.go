package main

import "fmt"

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

func main() {
	fmt.Println("SUM(1,5) =:", summate(1,5))
	fmt.Println("PROD(1,10) =:", product(1,10))
	fmt.Println("8! =:", factorial(8))
	fmt.Println("8C2 =:", combinations(8,3))
	fmt.Println("8P2 =:", permutations(8,3))
}
