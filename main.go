package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Starting pipeline simulation")

    // Simulating pipeline steps
    steps := []string{"Checkout code", "Set up Go", "Install dependencies", "Build application", "Run application"}
    for _, step := range steps {
        fmt.Printf("Executing step: %s\n", step)
        time.Sleep(time.Second) // Simulate some work
        fmt.Printf("Completed step: %s\n", step)
    }

    fmt.Println("Pipeline simulation completed")
}
