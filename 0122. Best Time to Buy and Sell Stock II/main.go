package main

func maxProfit(prices []int) int {
	max := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}

	}

	return max
}

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	maxProfit(a)
}
