package fib

// FibonacciIterative calculates the first n Fibonacci numbers using an iterative approach.
// It returns a slice of integers representing the sequence.
//
// The Fibonacci sequence starts with 0 and 1. Each subsequent number is the sum
// of the two preceding ones.
//
// Parameters:
//   n (int): The number of Fibonacci numbers to generate. Must be non-negative.
//
// Returns:
//   []int: A slice containing the first n Fibonacci numbers.
//
// Example:
//   FibonacciIterative(0) => []
//   FibonacciIterative(1) => [0]
//   FibonacciIterative(5) => [0, 1, 1, 2, 3]
func FibonacciIterative(n int) []int {
	if n < 0 {
		return []int{} // Or handle as an error in later phases
	}
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	fibSequence := make([]int, n)
	fibSequence[0] = 0
	fibSequence[1] = 1

	for i := 2; i < n; i++ {
		fibSequence[i] = fibSequence[i-1] + fibSequence[i-2]
	}

	return fibSequence
}
