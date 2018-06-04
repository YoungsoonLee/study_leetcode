package main

import (
	"fmt"
	"strconv"
)

// O(n)
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return true
	}

	j := len(s) - 1
	for i, c := range s {
		if c != rune(s[j]) {
			return false
		}
		if i == j {
			return true
		}
		j--
	}

	return true
}

// O(logN), because revert half
func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	rn := 0
	for x > rn {
		//fmt.Println(x%10, rn*10)
		rn = rn*10 + x%10
		x /= 10
	}
	return x == rn || x == rn/10
}

func main() {
	fmt.Println(isPalindrome2(1221))
}
