package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
		var aSize, input, dest, src, size int
		var array []int

		fmt.Scanf("%d", &aSize)

		array = make([]int, 0)

		for i := 0; i < aSize; i++ {
			fmt.Scan(&input)
			array = append(array, input)
		}
		fmt.Scanf("%d %d %d", &dest, &src, &size)
	*/
	/*
		array := []int{0, 255, 123, 12, 2, 4, 12, 4, 55, 2}

		dest := 5
		src := 0
		size := 3

		memcpy(array, dest, src, size)
	*/

	/*
		q := make([]int, 0)
		min := 1<<63 - 1
		fmt.Println(min)
		add(5, q, &min)
	*/

}

func add(n int, q []int, m *int) {
	q = append(q, n)
	saveMin(q, m)
}

func saveMin(q []int, m *int) {
	for i := 1; i < len(q)-1; i++ {
		*m = getMin(q[i-1], q[1])
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min(m *int) {
	fmt.Println(m)
}

func remove(q []int) {
	x := q[0]
	// Discard top element
	q = q[1:]
	fmt.Println(x)
}

func memcpy(array []int, dest, src, size int) {
	//fmt.Scanf("%d %d %d", &dest, &src, &size)
	//fmt.Println(aSize)
	fmt.Println(array)
	fmt.Println(dest, src, size)

	//result := ""

	if size <= 0 {
		printString(array)

	} else {
		//array = append(array[:dest], append(array[src:size], array[dest+size:]...)...)
		temp := array[dest+size:]
		//fmt.Println(temp)
		array = append(array[:dest], array[src:size]...)
		array = append(array[:dest+size], temp...)
		printString(array)
	}

}

func printString(a []int) {
	//fmt.Println("len: ", len(a))
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	fmt.Println(strings.Join(b, " "))
}

/*
func getArea(width, height int) {
	if width > 1024 || height > 102 {
		fmt.Println("overflow input. you should input 0 ~ 1024")
	}

	fmt.Println("Area is : ", width*height)
}
*/
