package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0

	for i < n-1 {
		//fmt.Println(i, bits[i])
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}

	fmt.Println(i, n-1)
	return i == n-1 // !!!!!
}

func main() {
	a := []int{1, 0, 0}
	fmt.Println(isOneBitCharacter(a))
}
