package main

import (
	"fmt"
	"sort"
)

func readBinaryWatch(num int) []string {
	res := make([]string, 0, 8)
	leds := make([]bool, 10)

	//fmt.Println(res)
	//fmt.Println(leds)

	var dfs func(int, int)
	dfs = func(idx, n int) {
		var h, m int
		if n == 0 {
			m, h = get(leds[:6]), get(leds[6:])
			if h < 12 && m < 60 {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
			return
		}

		for i := idx; i < len(leds)-n+1; i++ {
			leds[i] = true
			dfs(i+1, n-1)
		}
	}

	dfs(0, num)

	sort.Strings(res)
	fmt.Println("res: ", res)
	return res
}

var bs = []int{1, 2, 4, 8, 16, 32}

func get(leds []bool) int {
	var sum int
	size := len(leds)
	for i := 0; i < size; i++ {
		if leds[i] {
			sum += bs[i]
		}
	}
	//fmt.Println("sum: ", sum)
	return sum
}

func main() {
	a := 1
	readBinaryWatch(a)
}
