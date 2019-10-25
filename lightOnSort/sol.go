package sol

// you can also use imports, for example:
import (
	"sort"
) // import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// func main() {
// 	arr := [][]int{
// 		{1, 2, 3, 4, 5},
// 		{9, 1, 6, 7, 2, 3, 4, 8, 5},
// 		{3, 1, 2, 5, 4},
// 		{7, 4, 1, 2, 3, 6, 9, 8, 5},
// 	}
// 	for _, a := range arr {
// 		fmt.Println(a, sol(a))
// 	}

// }

func Sol(A []int) int {
	if len(A) == 0 {
		return 0
	}
	result := []int{}
	findNumber := 1
	for i := 1; i <= len(A); i++ {
		if A[i-1] == findNumber {
			result = append(result, A[i-1])
			findNumber++
		}
	}

	if A[len(A)-1] > findNumber {
		return len(result) + 1
	}

	return len(result)
}

func Solution(A []int) int {
	light := 0
	r := []int{}
	c := []int{}
	lastm := 0
	for i := 0; i < len(A); i++ {
		m := 0

		r = append(r, A[i])
		c = append(c, i+1)
		sort.Ints(r)
		for j := 0; j < len(r); j++ {
			if c[j] == r[j] {
				m++
			}
		}
		if m > lastm {
			light++
			lastm = m
		}

	}
	return light

}
