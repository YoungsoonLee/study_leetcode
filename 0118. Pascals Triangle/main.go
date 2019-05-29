package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{}

	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, getNext(res[i-1]))
	}

	fmt.Println("final: ", res)
	return res
}

func getNext(p []int) []int {
	r := make([]int, 1, len(p)+1)
	fmt.Println(r)
	r = append(r, p...)
	fmt.Println(r)

	for i := 0; i < len(r)-1; i++ {
		r[i] = r[i] + r[i+1]
		fmt.Println(r)
	}

	return r
}

func main() {
	n := 5
	generate(n)
}
