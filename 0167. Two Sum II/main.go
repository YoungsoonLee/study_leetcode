package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	re := make([]int, 0)
	hs := make(map[int]int)
	for i, v := range numbers {
		hs[v] = i
	}

	fmt.Println(hs)

	for i := range numbers {
		compl := target - numbers[i]
		if _, ok := hs[compl]; ok {
			if hs[compl] != i { // 이게 없으면 모두 결과로 나온다.
				re = append(re, i+1)
				re = append(re, hs[compl]+1)
				return re
			}
		}
	}
	return re

}

func twoSum2(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	fmt.Println(m)
	for i, n := range numbers {
		if m[target-n] != 0 {
			fmt.Println("if: ", m[target-n])
			return []int{m[target-n], i + 1}
		}
		m[n] = i + 1 // !!! 왜 i 는 0부터 시작
	}
	return nil
}

func main() {
	n := []int{2, 3, 4}
	t := 6

	fmt.Println(twoSum2(n, t))

}
