package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	tmp := make(map[int][]string, len(strs)/2)
	for _, s := range strs {
		c := encode(s)
		tmp[c] = append(tmp[c], s)
	}

	fmt.Println(tmp)

	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}
	fmt.Println(res)
	return res
	/*
		m := make(map[byte]bool, 0)
		var r [][]string

		for i := 0; i < len(strs); i++ {
			s := strs[i]
			c := 0
			f := make([]byte, 0)

			for j := 0; j < len(s); j++ {

				if i == 0 || m[s[j]] != true {
					m[s[j]] = true
				} else {
					f = append(f, s[j])
					c++
				}

			}
			fmt.Println(c, f)

			if c == 3 {
				r = append(r, string(f[:len(f)]))
			}

		}

		fmt.Println(r)
		return [][]string{{}, {}}
	*/
}

// 프라임은 A에서 Z에 해당합니다. 영어로 나타날 확률이 높을수록 선택된 소수를 더 작게 만듭니다.
var prime = []int{5, 71, 37, 29, 2, 53, 59, 19, 11, 83, 79, 31, 43, 13, 7, 67, 97, 23, 17, 3, 41, 73, 47, 89, 61, 101}

func encode(s string) int {
	res := 1
	for i := range s {
		fmt.Println(s, i, s[i]-'a', prime[s[i]-'a'])
		res *= prime[s[i]-'a']
	}
	return res
}

func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(a)
}
