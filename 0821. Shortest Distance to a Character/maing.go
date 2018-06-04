package main

import "fmt"

func shortestToChar(S string, C byte) []int {
	cIndex := make([]int, 0)
	// save C's index
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			cIndex = append(cIndex, i)
		}
	}

	fmt.Println(cIndex)

	ret := make([]int, 0, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] != C {
			//sd :=
			ret = append(ret, shortestDist(i, &cIndex))
		} else {
			ret = append(ret, 0)
		}
	}
	fmt.Println(ret)
	return ret
}

func shortestDist(i int, cI *[]int) int {
	sh := 63 << 1
	for _, ci := range *cI {
		sh = min(sh, abs(i-ci))
	}
	return sh
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "loveleetcode"
	c := byte('e')
	shortestToChar(s, c)
}
