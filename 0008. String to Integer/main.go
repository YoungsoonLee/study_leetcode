package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func myAtoi_uu(str string) int {
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}

	sign, x := getSign(s)
	x = trim(x)
	return convert(sign, x)
}

func convert(sign int, s string) int {
	base := 1 * sign
	res := 0
	yes := false

	for i := len(s) - 1; i >= 0; i-- {
		res, yes = isOverflow(res + (int(s[i])-48)*base)
		if yes {
			return res
		}

		base *= 10

		//fmt.Println(res)
	}

	return res
}

func isOverflow(i int) (int, bool) {
	switch {
	case i > math.MaxInt32:
		return math.MaxInt32, true
	case i < math.MinInt32:
		return math.MinInt32, true
	default:
		return i, false
	}
}

func trim(s string) string {
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return s[:i]
		}
	}
	return s
}

func getSign(s string) (int, string) {
	sign := 1
	switch s[0] {
	case '-':
		s = s[1:]
		sign = -1.0
	case '+':
		s = s[1:]
	default:
	}

	return sign, s
}

func myAtoi(str string) int {

	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	if len(str) == 1 && (str[0] == '-' || str[0] == '+') {
		return 0
	}

	//fmt.Println("[" + str + "]")
	//fmt.Println(isNumber(str))

	if !isNumber(str) {
		return 0
	}

	rt := getDigit(str)
	fmt.Println(rt, math.MinInt32)

	if rt < math.MinInt32 {
		return math.MinInt32
	}

	if rt > math.MaxInt32 {
		return math.MaxInt32
	}

	return rt

}

func getDigit(s string) int {
	//fmt.Println("getDigit:[" + s + "]")

	d := make([]string, 0)
	min := false

	if string(s[0]) == "-" {
		min = true
		s = s[1:]
	}

	if string(s[0]) == "+" {
		s = s[1:]
	}

	fmt.Println("getDigit:[" + s + "]")
	for _, c := range s {
		//fmt.Println(c, isNumber(string(c)))
		//if isNumber(string(c)) {
		if c >= '0' && c <= '9' {
			d = append(d, string(c))
		} else {
			break
		}
	}

	fmt.Println("getDigit: ", d)

	ret, _ := strconv.Atoi(strings.Join(d, ""))
	if min {
		return -1 * ret
	}
	return ret

}

func isNumber(s string) bool {
	if string(s[0]) == "-" || string(s[0]) == "+" {
		if s[1] >= '0' && s[1] <= '9' {
			return true
		}
	} else {
		if s[0] >= '0' && s[0] <= '9' {
			return true
		}
	}

	return false
}

func main() {
	s := "   -42"
	fmt.Println(myAtoi(s))
}
