package main

import (
	"fmt"
	"sort"
)

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

func main() {
	p := []int{1, 2, 3, 4, 5}
	maxProfit(p)
}
