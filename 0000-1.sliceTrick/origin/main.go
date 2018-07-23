package main

import (
	"fmt"
	"math/rand"
)

func ex_append(a, b []int) {
	a = append(a, b...)
	fmt.Println(a)
}

func ex_copy(a []int) {
	b := make([]int, len(a))
	copy(a, a)

	// or
	// b = append([]T(nil), a...)

	fmt.Println("copy: ", b)
}

func ex_cut(a []int) {
	/* this is memory leak
	a = append(a[:2], a[4:]...)
	fmt.Println(a)
	*/
	copy(a[2:], a[3:])
	fmt.Println("cut: ", a)
	fmt.Println("cut-1: ", len(a))
	for k, n := len(a)-3+2, len(a); k < n; k++ {
		a[k] = 0 // or the zero value of T
	}
	fmt.Println("cut-2: ", a)
	a = a[:len(a)-3+2]
	fmt.Println("cut-3: ", a)

}

func ex_delete(a []int) {
	/* this is memory leak
	a = append(a[:2], a[3:]...)
	fmt.Println("delete: ", a)
	*/

	i := 2
	copy(a[i:], a[i+1:])
	fmt.Println("delete1: ", a)
	a[len(a)-1] = 0 // or the zero value of T
	a = a[:len(a)-1]
	fmt.Println("delete2: ", a)

	/* without preserving order
	i := 2
	a[i] = a[len(a)-1]
	fmt.Println("delete1: ", a)
	a[len(a)-1] = 0
	fmt.Println("delete2: ", a)
	a = a[:len(a)-1]
	fmt.Println("delete3: ", a)
	*/
}

func ex_expand(a []int) {
	i := 2
	j := 3
	a = append(a[:i], append(make([]int, j), a[i:]...)...) // !!!
	fmt.Println("expand: ", a)
}

func ex_extend(a []int) { // append end of slice
	j := 3
	a = append(a, make([]int, j)...)
	fmt.Println("extend: ", a)
}

func ex_insert(a []int) {
	i := 2
	a = append(a[:i], append([]int{99}, a[i:]...)...)
	fmt.Println("insert: ", a)
}

func ex_insert2(a []int) {
	i := 2
	a = append(a, 0)
	fmt.Println("insert2-1: ", a)

	//fmt.Println(a[i+1:])
	//fmt.Println(a[i:])

	copy(a[i+1:], a[i:]) // !!! : 뒤는 갯수 ~!!!!!!!!!

	//fmt.Println(a[i+1:])
	//fmt.Println(a[i:])

	fmt.Println("insert2-2: ", a)

	a[i] = 99
	fmt.Println("insert2-3: ", a)
}

func ex_InsertVector(a, b []int) {
	i := 2
	a = append(a[:i], append(b, a[i:]...)...)
	fmt.Println("invert: ", a)
}

func ex_popshift(a []int) {
	var x int
	x, a = a[0], a[1:]
	fmt.Println("PopAndShift: ", x, a)
}

func ex_popback(a []int) {
	var x int
	x, a = a[len(a)-1], a[:len(a)-1]
	fmt.Println("PopAndBack: ", x, a)
}

func ex_Push_FrontUnshift(a []int) {
	a = append([]int{99}, a...)
	fmt.Println("Push_FrontUnshift: ", a)
}

func ex_filtering(a []int) {
	b := a[:0]
	fmt.Println("b-1: ", b)
	for _, x := range a {
		//if f(x) {
		if x == 5 {
			b = append(b, x)
		}
	}
	fmt.Println("b-2: ", b)
	fmt.Println(a, b)
}

// !!!!
func ex_reverse(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		fmt.Println("reverse i: ", i)
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println("reverse: ", a)
}

func ex_reverse2(a []int) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	fmt.Println("reverse:2 ", a)
}

func ex_shuffling(a []int) {
	// Fisher–Yates algorithm:
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
	ex_append(a, b)
	ex_copy(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_cut(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_delete(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_expand(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_extend(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_insert(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_insert2(a)
	a = []int{1, 2, 3}
	b = []int{4, 5, 6}
	ex_InsertVector(a, b)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_popshift(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_popback(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_Push_FrontUnshift(a)
	a = []int{1, 2, 3, 4, 5, 6}
	ex_filtering(a)
	a = []int{1, 2, 3, 4, 5, 6, 7}
	ex_reverse(a)
	a = []int{1, 2, 3, 4, 5, 6, 7}
	ex_reverse2(a)
	a = []int{1, 2, 3, 4, 5, 6, 7}
	ex_shuffling(a)
}
