package main

import "fmt"

func characterReplacement(s string, k int) int {
	base := s[0]
	//count := 1
	t := ""
	for i := 1; i < len(s); i++ {
		if s[i] == base {
			continue
		} else {
			if k > 0 {
				t = s[:i] + string(base) + s[i:]
				//count++
				k--
			}
		}
	

	fmt.Println(t)
	return 0
}

func main() {
	s := "ABAB"
	k := 2
	characterReplacement(s, k)
}
