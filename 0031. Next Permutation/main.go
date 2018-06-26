package main

import "fmt"

func nextPermutation(a []int) {
	left := len(a) - 2 // why 2 ??

	fmt.Println("1: ", left, a[left], a[left+1])

	for 0 <= left && a[left] >= a[left+1] {
		left--
	}

	//이 시점에서 [left + 1 :]은 감소합니다
	reverse(a, left+1)
	fmt.Println("2: ", a)

	if left == -1 {
		return
	}

	//이 시점에서 [left + 1 :]은 증가하는 순서입니다
	right := search(a, left+1, a[left])
	fmt.Println("3: ", right, a)
	a[left], a[right] = a[right], a[left]
	fmt.Println("4: ", a)
}

func reverse(a []int, l int) {
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

// [l :] 행에있는> target의 최소값의 인덱스 번호를 반환합니다.
// a [l :]은 증가하는 순서입니다.
func search(a []int, l, target int) int {
	r := len(a) - 1
	l--
	for l+1 < r {
		mid := (l + r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func main() {
	a := []int{1, 2, 3}
	nextPermutation(a)
}
