package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0, right-left+1)
	for i := left; i <= right; i++ {
		if isSelfDriving(i) {
			res = append(res, i)
		}
	}
	fmt.Println(res)
	return res
}

func isSelfDriving(n int) bool {
	t := n
	for t > 0 {
		d := t % 10
		t /= 10
		if d == 0 || n%d != 0 {
			return false
		}
	}
	return true
}

func main() {
	l := 1
	r := 22
	selfDividingNumbers(l, r)
}
