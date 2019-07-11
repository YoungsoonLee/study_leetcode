package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	//stmap := make(map[byte]byte)
	//tsmap := make(map[byte]byte)

	stmap := make([]byte, 256)
	tsmap := make([]byte, 256)

	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]

		if stmap[a] != 0 {
			if stmap[a] != b {
				return false
			}

		} else if tsmap[b] != 0 {
			if tsmap[b] != a {
				return false
			}
		} else {
			stmap[a] = b
			tsmap[b] = a
		}
	}

	fmt.Println(stmap)
	fmt.Println(tsmap)

	return true

	/*
		if len(s) != len(t) {
			return false
		}

		m := make(map[byte]byte)

		for i := 0; i < len(s); i++ {
			if _, exists := m[s[i]]; exists {
				if m[s[i]] != t[i] {
					return false
				}
			}
			m[s[i]] = t[i]
		}
		fmt.Println(m)
		return true
	*/
}

/*
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
*/

func main() {
	s1 := "paper"
	s2 := "title"
	fmt.Println(isIsomorphic(s1, s2))
}
