package main

import "fmt"

func longestValidParentheses(s string) int {
	var left, max, temp int
	record := make([]int, len(s))
	fmt.Println("record1 : ", record)

	// 통계 기록
	for i, b := range s {
		if b == '(' {
			left++
		} else if left > 0 {
			left--
			record[i] = 2
		}
	}
	fmt.Println("record2 : ", record)

	// 레코드 수정
	for i := 0; i < len(record); i++ {
		if record[i] == 2 {
			j := i - 1
			for record[j] != 0 {
				j--
			}
			record[i], record[j] = 1, 1
		}
	}
	fmt.Println("record3 : ", record)

	// 통계
	for _, r := range record {
		if r == 0 {
			temp = 0
			continue
		}
		temp++
		if temp > max {
			max = temp

		}
	}

	fmt.Println("max: ", max)
	return max
}

func main() {
	s := "()(())" // "()(())"
	longestValidParentheses(s)
}
