package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make([]int, 256) // charaters has 256 ?
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		//fmt.Println(m1)
		//fmt.Println(m2)
		if m1[int(s[i])] != m2[int(t[i])] {
			//fmt.Println(m1)
			//fmt.Println(m2)
			return false
		}

		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1

	}
	fmt.Println(m1)
	fmt.Println(m2)
	return true
}

func main() {
	s1 := "aba"
	s2 := "bab"
	fmt.Println(isIsomorphic(s1, s2))
}
