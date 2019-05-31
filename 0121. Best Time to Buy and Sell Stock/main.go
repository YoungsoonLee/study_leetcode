package main

import "fmt"

// my solution - O(n^2)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	if len(prices) == 1 {
		return 0
	}
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				if max < prices[j]-prices[i] {
					max = prices[j] - prices[i]
				}
			}
		}
	}
	fmt.Println(max)
	return max
}

func maxProfit2(prices []int) int {
	max := 0
	temp := 0
	//fmt.Println("oo")
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		fmt.Println(temp)
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}
	//fmt.Println(max)
	return max
}

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	maxProfit2(a)
}
