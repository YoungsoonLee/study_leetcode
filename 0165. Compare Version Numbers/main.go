package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1s := conv(version1)
	v2s := conv(version2)

	fmt.Println(v1s)
	fmt.Println(v2s)

	if len(v1s) != len(v2s) {
		v1s, v2s = toSameLen(v1s, v2s)
	}

	fmt.Println(v1s)
	fmt.Println(v2s)

	for i := 0; i < len(v1s); i++ {
		if v1s[i] < v2s[i] {
			return -1
		} else if v1s[i] > v2s[i] {
			return 1
		}
	}
	return 0
}

func conv(version string) []int {
	vs := strings.Split(version, ".")
	fmt.Println("vs: ", vs)
	res := make([]int, len(vs))

	for i, v := range vs {
		res[i], _ = strconv.Atoi(v)
	}

	return res
}

func toSameLen(a1, a2 []int) ([]int, []int) {
	if len(a1) > len(a2) {
		res := make([]int, len(a1)) // !!! 크기는 큰거로
		copy(res, a2)               // !!! 크기 비워진것은 초기화 0
		return a1, res
	}

	r2, r1 := toSameLen(a2, a1)
	return r1, r2
}

func main() {
	s1 := "1.01"
	s2 := "1"
	fmt.Println(compareVersion(s1, s2))
}
