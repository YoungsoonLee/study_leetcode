package main

import "fmt"

func distributeCandies(candies []int) int {
	n := len(candies)
	fmt.Println(n)

	r := make(map[int]bool, n)
	for _, c := range candies {
		r[c] = true
	}

	fmt.Println(r)
	fmt.Println(min(len(r), n/2))
	return min(len(r), n/2) //!!!
}

/*
func distributeCandies(candies []int) int {
	sort.Ints(candies)
	ret := make([]int, 0)
	pu := 0
	j := 0

	for i := len(candies) - 1; i > 0; i-- {
		mx := max(candies[i], candies[i-1])
		if pu != mx {
			if j < len(candies)/2 {
				break
			} else {
				ret = append(ret, mx)
				pu = mx
				j++
			}

		}

	}

	return len(ret)
}
*/

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	candies := []int{1, 1, 2, 3}
	fmt.Println(distributeCandies(candies))
}
