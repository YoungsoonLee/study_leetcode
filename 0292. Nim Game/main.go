package main

import "fmt"

func canWinNim(n int) bool {
	if n > 0 && n <= 3 {
		return true
	}

	if n%3 == 0 {
		return false
	}

	return true
}

func canWinNim2(n int) bool {
	return n%4 != 0
}

func main() {
	a := 6
	fmt.Println(canWinNim(a))
}
