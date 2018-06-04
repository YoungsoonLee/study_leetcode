package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	// s의 "-"수를 센다.
	countDashs := strings.Count(s, "-")

	// s에서 문자와 숫자의 수를 센다.
	countAN := len(s) - countDashs

	if countAN == 0 {
		return ""
	}

	// '-'를 삭제하고 모든 문자를 대문자로 변경합니다.
	s = strings.Replace(s, "-", "", -1)
	s = strings.ToUpper(s)

	// (countAN + k-1) / k-1은 res에서 '-'의 수입니다
	res := make([]byte, (countAN+k-1)/k-1+countAN)

	i, j := len(res), len(s)
	for 0 <= j-k {
		copy(res[i-k:i], s[j-k:j])

		if 0 <= i-k-1 {
			res[i-k-1] = '-'
		}
		i -= k + 1
		j -= k
	}

	fmt.Println(res, s, j)

	// 불완전한 문자를 처리 합니다.
	if j > 0 {
		copy(res[:j], s[:j])
	}
	fmt.Println(res, s)
	return string(res)
}

func main() {
	s := "2-5g-3-J"
	k := 2
	fmt.Println(licenseKeyFormatting(s, k))
}
