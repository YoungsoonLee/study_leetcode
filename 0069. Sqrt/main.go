package main

import (
	"fmt"
	"math"
)

// x의 제곱근 구하기.
func mySqrt(x int) int {
	// 제곱근에 대한 Newton의 방법
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}

//binary search solution
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}

	left := 1
	right := math.MaxInt32

	fmt.Println("right: ", right)

	for {
		mid := left + (right-left)/2
		fmt.Println("mid: ", mid, " left: ", left, "right: ", right)

		if mid > x/mid {
			right = mid - 1
		} else {
			if mid+1 > x/(mid+1) {
				return mid
			}
			left = mid + 1
		}
	}
}

func main() {
	mySqrt2(8)
}
