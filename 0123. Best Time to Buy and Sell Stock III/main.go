package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	profits := []int{}
	temp := 0

	for i := 0; i < size; i++ {
		diff := prices[i] - prices[i-1]
		if temp*diff >= 0 {
			temp += diff
			continue
		}

		profits = append(profits, temp)
		temp = diff
	}
	profits = append(profits, temp)

	fmt.Println(profits)

	return 0
}

/* my solution but failed.
func maxProfit(prices []int) int {
	temp := []int{}

	for i := 0; i < len(prices)-2; i++ {
		for j := i + 1; j < len(prices)-1; j++ {
			if prices[j] > prices[i] {
				temp = append(temp, prices[j]-prices[i])
			}
		}
	}

	//fmt.Println(temp)
	if len(temp) == 0 {
		return 0
	}
	if len(temp) == 1 {
		return temp[0]
	}

	fmt.Println(temp)
	sort.Ints(temp)
	fmt.Println(temp)

	s := temp[len(temp)-2] + temp[len(temp)-1]

	println(s)

	return s
}
*/

func main() {
	p := []int{1, 2, 3, 4, 5}
	maxProfit(p)
}
