package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	remains, debts, start := 0, 0, 0

	for i, g := range gas {
		remains += g - cost[i]
		fmt.Println("1: ", i, g, cost[i], remains)

		if remains < 0 {
			// i + 1에서 다시 시작
			start = i + 1
			// 길을 따라 누락 된 기름의 양을 기록하십시오.
			debts += remains
			// 0으로 남음
			remains = 0

			fmt.Println("2: ", start, debts)
		}
	}

	fmt.Println("3: ", debts, remains)
	if debts+remains < 0 {
		// 이전 기간의 오일 량의 전액 상환과 같은 마지막 오일 잔량
		// 그런 다음 무릎을 마칠 수 없다.
		return -1
	}

	fmt.Println("4: ", start)
	return start
}

func main() {
	//g := []int{1, 2, 3, 4, 5}
	//c := []int{3, 4, 5, 1, 2}
	g := []int{2, 3, 4}
	c := []int{3, 4, 3}
	fmt.Println("r: ", canCompleteCircuit(g, c))
}
