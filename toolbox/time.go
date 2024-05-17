package main

import (
    "fmt"
    "time"
)

func tgif() {
    fmt.Println("Når Fredag?")
    today := time.Now().Weekday()

    // Omit the switch condition for equivalent functionality as if-elseif-else chains, but cleaner
    switch time.Friday {
        case today:
            fmt.Println("Idag")
        case today + 1:
            fmt.Println("I morgen.")
        case today + 2:
            fmt.Println("I to dager.")
        default:
            fmt.Println("I over to dager!")
    }
}

func main() {
    tgif()

    fmt.Println("Explosive device armed and ready...")
    defer fmt.Println("Kablooey!")
    
    // stacked deferred functions
    for i := 1; i <= 10; i++ {
        defer fmt.Printf("Detonation in T-%d\n", i)
    }
}

