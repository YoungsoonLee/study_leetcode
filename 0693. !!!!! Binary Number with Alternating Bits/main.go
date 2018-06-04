package main

import (
	"fmt"
)

func hasAlternatingBits_my(n int) bool {
	//fmt.Println(strconv.FormatInt(int64(n), 2))
	s := fmt.Sprintf("%b", n)
	fmt.Println(s)

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}
	return true
}

func hasAlternatingBits(n int) bool {
	std := n & 3 // 비트연산
	fmt.Println(std)
	//fmt.Println(fmt.Sprintf("%b", std))

	if std != 1 && std != 2 {
		return false
	}

	n >>= 2 // 비트 이동 2만큼
	fmt.Println(n)

	for n > 0 {
		if n&3 != std {
			return false
		}
		n >>= 2
	}

	return false
}

func main() {
	a := 10
	fmt.Println(hasAlternatingBits(a))
}
