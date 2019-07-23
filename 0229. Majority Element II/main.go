package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func majorityElement(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	e1, e2, c1, c2 := 0, 1, 0, 0
	for _, e := range nums {
		switch {
		case e == e1:
			c1++
		case e == e2:
			c2++
		case c1 == 0:
			e1 = e
			c1 = 1
		case c2 == 0:
			e2 = e
			c2 = 1
		default:
			c1--
			c2--
		}
	}

	res := []int{}

	if maj(nums, e2) {
		res = append(res, e2)
	}

	if maj(nums, e1) {
		res = append(res, e1)
	}
	return res
}

func maj(a []int, n int) bool {
	c := 0
	for _, e := range a {
		if e == n {
			c++
		}
	}
	return c > (len(a) / 3)
}

func main() {

	b := make([]byte, 6) //equals 8 charachters
	rand.Read(b)
	s := hex.EncodeToString(b)
	fmt.Println(s)
}
