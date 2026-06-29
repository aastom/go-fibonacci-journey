package fib

import (
	"reflect"
	"testing"
)

func TestFibonacciIterative(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{name: "n=0", n: 0, want: []int{}},
		{name: "n=1", n: 1, want: []int{0}},
		{name: "n=2", n: 2, want: []int{0, 1}},
		{name: "n=5", n: 5, want: []int{0, 1, 1, 2, 3}},
		{name: "n=10", n: 10, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{name: "n_negative", n: -5, want: []int{}}, // Should return empty for negative input
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciIterative(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FibonacciIterative(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
