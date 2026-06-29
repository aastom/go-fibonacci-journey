package main

import (
	"fmt"
	"go-fibonacci-journey/pkg/fib"
)

// This is Phase 1 of the Go Fibonacci Learning Journey.
// Goal: Introduce basic Go syntax, function definition, loop structures,
// and standard output. Provide a quick win with a runnable program.
//
// How to run:
// go run ./cmd/01_basic_iterative
//
// Expected Output:
// Fibonacci sequence (first 10 numbers): [0 1 1 2 3 5 8 13 21 34]
func main() {
	// Define how many Fibonacci numbers to generate for this phase
	n := 10

	// Calculate the Fibonacci sequence using the iterative function from the fib package
	sequence := fib.FibonacciIterative(n)

	// Print the generated sequence to standard output
	fmt.Printf("Fibonacci sequence (first %d numbers): %v\n", n, sequence)
}
