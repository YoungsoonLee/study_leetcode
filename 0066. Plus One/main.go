package main

import "fmt"

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

	/* faild when a big number
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
	*/
}

func main() {
	i := []int{1, 2, 3}
	fmt.Println(plusOne(i))
}
