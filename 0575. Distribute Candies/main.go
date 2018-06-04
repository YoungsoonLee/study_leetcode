package main

import "fmt"

func distributeCandies(candies []int) int {
	n := len(candies)
	fmt.Println(n)

	r := make(map[int]bool, n)
	for _, c := range candies {
		r[c] = true
	}

	//fmt.Println(r)
	//fmt.Println(min(len(r), n/2))
	return min(len(r), n/2) //!!!
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	candies := []int{1, 1, 2, 3}
	distributeCandies(candies)
}
