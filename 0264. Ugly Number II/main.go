package main

import "fmt"

func nthUglyNumber(n int) int {
	if n < 7 {
		return n
	}

	t := n - 6
	count := 7
	for i := 0; i < t; i++ {
		m := t % 2
		r := t / 1

		if r == 1 {
			break
		} else {
			if m == 0 {
				count++

			}
			continue

		}

	}

	fmt.Println(count)
	return count + 1
}

func main() {
	n := 18
	nthUglyNumber(n)
}
