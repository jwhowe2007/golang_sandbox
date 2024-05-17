package main

import (
    "fmt"
    "runtime"
)

func mySystem() {
    fmt.Print("Go is running on: ")

    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux! Excellent!!")
    default:
        // other
        fmt.Print("%s.\n", os)
    }
}

func pointerResearch() {
    i, j := 42, 2701
    p := &i
    fmt.Println(*p)

    *p = 21
    fmt.Println(i)

    p = &j
    *p /= 37
    fmt.Println(j)
}

func main() {
    fmt.Println("go" + "lang") // String concat

    fmt.Println("1+1 =", 1+1) // Addition, print args
    fmt.Println("7.0/3.0 =", 7.0/3.0) // Floating point arithmetic

    fmt.Println("True && False:", true && false)
    fmt.Println("True || False:", true || false)
    fmt.Println("Not True (false):", !true)
    fmt.Println("Inference Rule (if p then q) (FALSE):", !true || false)
    fmt.Println("Inference Rule (if p then q) (TRUE):", !false || true)

    mySystem()
    pointerResearch()
}
