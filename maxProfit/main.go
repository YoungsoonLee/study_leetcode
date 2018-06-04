package main

import "fmt"

func maxProfit(data []int) int {
	//size := len(data)
	maxProfit := 0
	minPrice := data[0]

	for _, v := range data[1:] {
		p := v - minPrice
		fmt.Println(p)
		if p > maxProfit {
			maxProfit = p
		}
		if p < minPrice {
			minPrice = v
		}
	}

	return maxProfit
}

func main() {
	a := []int{10300, 9600, 9800, 8200, 7800, 8300, 9500, 9800, 10200, 9500}
	fmt.Println(maxProfit(a))
}
