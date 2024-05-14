package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") // String concat

	fmt.Println("1+1 =", 1+1) // Addition, print args
	fmt.Println("7.0/3.0 =", 7.0/3.0) // Floating point arithmetic

	fmt.Println("True && False:", true && false)
	fmt.Println("True || False:", true || false)
	fmt.Println("Not True (false):", !true)
	fmt.Println("Inference Rule (if p then q) (FALSE):", !true || false)
	fmt.Println("Inference Rule (if p then q) (TRUE):", !false || true)
}