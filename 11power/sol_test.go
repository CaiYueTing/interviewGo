package sol

import "testing"

func TestSolver(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{1, Tsol(1)},
		{2, Tsol(2)},
		{3, Tsol(3)},
		{5, Tsol(5)},
		{8, Tsol(8)},
		{18, 4},
		{19, 4},
		{20, 1},
		{40, 3},
		{50, 6},
		{60, 6},
	}
	for _, test := range tests {
		if got := Sol(test.input); got != test.want {
			t.Error("wrong", test.input, got, test.want)
		}

	}
}
