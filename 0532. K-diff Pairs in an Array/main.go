package main

import "fmt"

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	record := make(map[int]int)
	for _, num := range nums {
		record[num]++
	}

	fmt.Println(record)

	ans := 0

	if k == 0 {
		for _, count := range record {
			if count > 1 {
				ans++
			}
		}
		return ans
	}

	// !!!
	for n := range record {
		fmt.Println(n)
		if record[n-k] > 0 {
			//fmt.Println(n, n-k, record[n-k])
			ans++
		}
	}

	fmt.Println("ans: ", ans)
	return ans
}

/*
func findPairs(nums []int, k int) int {
	if k == 0 {
		return len(nums)
	}
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	fmt.Println(m)
	count := 0
	ad := make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n+k]; ok {
			count++

			if _, ok := ad[n+k]; ok {
				delete(m, n+k)
			} else {
				ad[n] = true
			}
		}
	}

	fmt.Println(count)
	return count
}
*/

func main() {
	a := []int{3, 1, 4, 1, 5}
	k := 2
	findPairs(a, k)
}
