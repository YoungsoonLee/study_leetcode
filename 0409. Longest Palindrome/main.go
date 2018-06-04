package main

import "fmt"

/*
func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}

	if len(s) == 2 {
		return 2
	}

	hm := make(map[string]int)
	for _, c := range s {
		hm[string(c)]++
	}

	fmt.Println(hm)
	ret := 0
	for _, v := range hm {
		if v > 1 {
			ret += v
		}
	}

	return ret + 1
}
*/
func longestPalindrome(s string) int {

	res := 0
	a := [123]int{} // The ASCII code of 'z' is 122
	for i := range s {
		//fmt.Println(s[i], i)
		a[s[i]]++
	}

	// hasOdd는 중간에 놓일 수있는 홀수 개의 요소가 있음을 나타냅니다.
	hasOdd := 0
	for i := range a {
		if a[i] == 0 {
			continue
		}
		fmt.Println(a[i]&1, a[i])
		if a[i]&1 == 0 { // 1일때만 1을 만들고, 나머진 다 0으로 만든다.
			res += a[i]
		} else {
			res += a[i] - 1
			hasOdd = 1
		}
	}
	//fmt.Println(a)
	return res + hasOdd
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
