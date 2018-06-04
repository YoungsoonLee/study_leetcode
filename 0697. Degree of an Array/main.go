package main

import "fmt"

func findShortestSubArray(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	first := make(map[int]int, size)
	count := make(map[int]int, size)

	maxCount := 1
	minLen := size

	for i, n := range nums {
		count[n]++
		if count[n] == 1 {
			first[n] = i
		} else {
			l := i - first[n] + 1
			if maxCount < count[n] || (maxCount == count[n] && minLen > l) {
				maxCount = count[n]
				minLen = l
			}
		}

		fmt.Println(i, n, count, first, maxCount, minLen)
	}

	if len(count) == size {
		return 1
	}

	return minLen
}

func main() {
	a := []int{1, 2, 2, 3, 1, 4, 2}
	findShortestSubArray(a)
}
