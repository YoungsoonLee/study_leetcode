package main

import (
	"fmt"
	"sort"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	i := sort.Search(n, func(i int) bool {
		fmt.Println(i, n)
		return target < letters[i]
	})
	//fmt.Println(i, n)
	return letters[i%n]
}

func main() {
	a := []byte{'c', 'f', 'j'}
	t := byte('k')
	fmt.Println(nextGreatestLetter(a, t))
}
