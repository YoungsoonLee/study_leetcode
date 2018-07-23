package main

import (
	"fmt"
	"math/rand"
)

func exAppend(a, b []int) {
	a = append(a, b...)
	fmt.Println(a)
	//fmt.Println(b)
}

func exCopy(a []int) {
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println("copy: ", b)
}

func exCut(a []int) {
	/* memory leak
	a = append(a[:2], a[4:]...)
	fmt.Println("cut: ", a)
	*/
	copy(a[:2], a[3:])
	fmt.Println("cut: ", a)
	fmt.Println("cut-1: ", len(a))
	for k, n := len(a)-3+2, len(a); k < n; k++ {
		a[k] = 0 // or the zero value of T
	}
	fmt.Println("cut-2: ", a)
	a = a[:len(a)-3+2]
	fmt.Println("cut-3: ", a)
}

func exDelete(a []int) {
	/* */
}

func exExpand(a []int) {
	i := 2
	j := 3
	a = append(a[:i], append(make([]int, j), a[i:]...)...)
	fmt.Println("expand: ", a)
}

// append end of slice
func exExtend(a []int) {
	j := 3
	a = append(a, make([]int, j)...)
	fmt.Println("extend: ", a)
}

func exInsert(a []int) {
	i := 2
	a = append(a[:i], append([]int{99}, a[i:]...)...)
	fmt.Println("insert: ", a)
}

func exInsert2(a []int) {
	i := 2
	a = append(a, 0)
	fmt.Println("insert2-1: ", a)

	fmt.Println(a[i+1:])
	fmt.Println(a[i:])

	copy(a[i+1:], a[i:]) // !!! : 뒤는 갯수 ~!!!!!!!!!
	fmt.Println(a)

	a[i] = 99
	fmt.Println("insert2-3: ", a)
}

func exInvert(a, b []int) {
	i := 2
	a = append(a[:i], append(b, a[i:]...)...)
	fmt.Println("invert: ", a)
}

func exPopShift(a []int) {
	var x int
	x, a = a[0], a[1:]
	fmt.Println("PopAndShift: ", x, a)
}

func exPopBack(a []int) {
	var x int
	x, a = a[len(a)-1], a[:len(a)-1]
	fmt.Println("PopAndBack: ", x, a)
}

func exPushFrontUnShift(a []int) {
	a = append([]int{99}, a...)
	fmt.Println("Push_FrontUnshift: ", a)
}

func exFiltering(a []int) {
	b := a[:0]
	fmt.Println("b-1: ", b)
	fmt.Println("b-1: ", a)
	for _, x := range a {
		if x == 5 {
			b = append(b, x)
		}
	}
	fmt.Println("b-2: ", b)
	fmt.Println(a, b)
}

// !!!
func exReverse(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		fmt.Println("reverse i: ", i)
		opp := len(a) - 1 - i
		fmt.Println("reverse opp: ", opp)
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println("reverse: ", a)
}

func exReverse2(a []int) {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	fmt.Println("reverse:2 ", a)
}

func exShuffling(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		fmt.Println("j: ", j)
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println("shuffling: ", a)
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	exAppend(a, b)
	exCopy(a)
	a = []int{1, 2, 3, 4, 5, 6}
	//exCut(a)
	exExpand(a)
	exExtend(a)
	exInsert(a)
	exInsert2(a)
	a = []int{1, 2, 3}
	b = []int{4, 5, 6}
	exInvert(a, b)
	a = []int{1, 2, 3, 4, 5, 6}
	exPopShift(a)
	exPopBack(a)
	exPushFrontUnShift(a)
	exFiltering(a)
	a = []int{1, 2, 3, 4, 5, 6}
	exReverse(a)
	exReverse2(a)
	exShuffling(a)
}
