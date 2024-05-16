package main

import (
    "fmt"
    "time"
)

func main() {
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