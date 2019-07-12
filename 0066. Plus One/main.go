package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			if digits[1] == 9 {
				digits[i] = 0
				carry = true
			} else {
				digits[i]++
				carry = false
			}
		} else {
			break
		}
	}

	if carry {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func ConvertInt(digits []int) int {
	no := 0
	for i := 0; i < len(digits); i++ {
		no *= 10
		no += digits[i]
	}
	return no
}

func ConvertIntArray(n int) []int {
	res := make([]int, 0)
	for n > 0 {
		d := n % 10
		res = append([]int{d}, res...)
		n /= 10
	}
	return res
}

/*
func plusOne(digits []int) []int {
	size := len(digits)
	for i := size - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}

	// 자리수 하나 더 만들기.. !!!
	if digits[0] == 0 {
		nn := make([]int, size+1, size+1)
		nn[0] = 1
		return nn
	}

	return digits

	-* faild when a big number
	s := ""
	for _, n := range digits {
		s = s + strconv.Itoa(n)
	}

	fmt.Println(s)
	sn, _ := strconv.Atoi(s)

	r := strconv.Itoa(sn + 1)
	fmt.Println(r)
	//fmt.Println(r[0])

	result := make([]int, 0)
	for _, c := range r {
		t, _ := strconv.Atoi(string(c))
		result = append(result, t)
	}
	fmt.Println(result)
	return digits
	*-
}
*/

func main() {
	i := []int{1, 2, 3}
	fmt.Println(plusOne(i))
}
