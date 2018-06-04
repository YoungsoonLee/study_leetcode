package main

import (
	"fmt"
)

func addStrings(s1 string, s2 string) string {
	/*
		carry := 0
		ret := make([]int, 0)
		i, j := len(s1), len(s2)
		for i > 0 && j > 0 {
			s := int(s1[i-1]+s2[j-1]) + carry
			if s > 10 {
				carry = int(s / 10)
				ret = append(ret, int(s%10))
			} else {
				ret = append(ret, int(s))
			}
			i--
			j--
		}

		fmt.Println(ret)
		return ""
	*/

	// 确保 n1 <= n2
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	n1, n2 := len(s1), len(s2)
	a1, a2 := []byte(s1), []byte(s2)

	//fmt.Println(n1, n2)
	//fmt.Println(a1, a2)
	carry := byte(0)
	// buf는 응답을 [] 바이트 형식으로 저장합니다.
	buf := make([]byte, n2+1)
	buf[0] = '1'

	i := 1
	for i <= n2 {
		// a1과 a2를 추가합니다.
		if i <= n1 {
			buf[n2+1-i] = a1[n1-1] - '0'
		}
		buf[n2+1-i] += a2[n2-i] + carry
		// carry 문제 처리
		if buf[n2+1-i] > '9' {
			buf[n2+1-i] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}

		i++
	}

	if carry == 1 {
		return string(buf)
	}
	return string(buf[1:])

}

func main() {
	n1 := "123"
	n2 := "456"
	fmt.Println(addStrings(n1, n2))
}
