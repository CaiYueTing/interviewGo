package sol

import (
	"strconv"
	"strings"
)

func Sol(input int) int {

	el := []int{1, 1}
	result := []int{}
	for i := 1; i < input; i++ {
		el = append([]int{0}, el...)
		for j := 0; j < len(el); j++ {
			if j+1 != len(el) {
				sum := el[j] + el[j+1]
				result = append(result, sum)
			} else {
				result = append(result, el[j])
			}
		}
		el = result
		result = []int{}
	}
	// fmt.Println(el)
	count := 0
	carry := 0
	for i := len(el) - 1; i >= 0; i-- {
		sum := el[i] + carry
		if sum > 9 {
			el[i] = (sum) % 10
			carry = sum / 10
		} else {
			el[i] = sum
			carry = 0
		}

		if el[i] == 1 {
			count++
		}
	}

	for carry != 0 {
		nel := carry % 10
		if nel == 1 {
			count++
		}
		el = append([]int{nel}, el...)
		carry = carry / 10
	}

	return count
}

func Esol(input int) int {
	el := []int{1, 1}
	result := []int{}
	for i := 1; i < input; i++ {
		el = append([]int{0}, el...)

		el = result
		result = []int{}
	}

	count := 0
	for _, v := range el {
		if v == 1 {
			count++
		}
	}
	return count
}

func Tsol(input int) int {
	if input < 1 {
		return 1
	}
	result := 1
	for i := 0; i < input; i++ {
		result = 11 * result
	}
	// fmt.Println(result)
	str := strconv.Itoa(result)
	arr := strings.Split(str, "")
	count := 0
	for _, v := range arr {
		if m, _ := strconv.Atoi(v); m == 1 {
			count++
		}
	}

	return count
}
