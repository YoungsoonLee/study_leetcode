package main

import (
	"fmt"
	"math"
)

// newton method
func mySqrt(x int) int {
	var r float64
	r = float64(x)

	for r*r > float64(x) {
		r = (r + float64(x)/r) / 2

		if int(r)%10000000 == 0 {
			break
		}
	}
	return int(r)
}

// use binary search
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	left := 1
	right := math.MaxInt32

	for {
		mid := left + (right-left)/2
		if mid > x/mid {
			right = mid - 1
		} else {
			if mid+1 > x/(mid+1) {
				return mid
			}
			left = mid + 1
		}

		fmt.Println(mid, left, right)
	}
}

func main() {
	a := 5
	fmt.Println(mySqrt2(a))
}
