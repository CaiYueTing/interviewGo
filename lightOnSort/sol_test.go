package sol

import (
	"testing"
)

func TestSolver(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{9, 1, 6, 7, 2, 3, 4, 8, 5}, 5},
		{[]int{3, 1, 2, 5, 4}, 3},
		{[]int{7, 4, 1, 2, 3, 6, 9, 8, 5}, 4},
	}

	for _, test := range tests {
		if got := Sol(test.input); got != test.want {
			t.Error("wrong", test.input, got)
		}

	}
}

func TestSolver1(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{9, 1, 6, 7, 2, 3, 4, 8, 5}, 5},
		{[]int{3, 1, 2, 5, 4}, 3},
		{[]int{7, 4, 1, 2, 3, 6, 9, 8, 5}, 4},
	}

	for _, test := range tests {
		if got := Solution(test.input); got != test.want {
			t.Error("wrong", test.input, got)
		}

	}
}

func BenchmarkSolver(b *testing.B) {
	// arr =[]int{9, 1, 6, 7, 2, 3, 4, 8, 5},
	for i := 0; i < b.N; i++ {
		Sol([]int{9, 1, 6, 7, 2, 3, 4, 8, 5})
	}
}
func BenchmarkSolver2(b *testing.B) {
	// arr =[]int{9, 1, 6, 7, 2, 3, 4, 8, 5},
	for i := 0; i < b.N; i++ {
		Solution([]int{9, 1, 6, 7, 2, 3, 4, 8, 5})
	}
}
