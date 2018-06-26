package main

import (
	"fmt"
	"math"
)

func divide(m, n int) int {
	// 누군가가 0을 제수로 취급하지 못하도록합니다.
	if n == 0 {
		return math.MaxInt32
	}

	signM, absM := analysis(m)
	signN, absN := analysis(n)

	res, _ := d(absM, absN, 1)
	if signM != signN {
		res = res - res - res
	}

	// 오버플로 검사
	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

// d는 m / n의 값을 계산하여 결과와 나머지를 반환합니다.
// m> = 0
// n> 0
// count == 1, 재귀 적 프로세스에서 초기 n의 수를 나타냅니다. count == 현재 n / 초기 n
func d(m, n, count int) (res, reminder int) {
	switch {
	case m < n:
		return 0, m
	case n <= m && m < n+n:
		return count, m - n
	default:
		res, reminder = d(m, n+n, count+count)
		if reminder >= n {
			return res + count, reminder - n
		}
		return
	}
}

func analysis(num int) (sign, abs int) {
	sign = 1
	abs = num
	if num < 0 {
		sign = -1
		abs = num - num - num
	}
	return
}

/*
func divide(dividend int, divisor int) int {
	re := 0
	absDividen, mDd := abs(dividend)
	absDivisor, mDr := abs(divisor)

	for absDividen >= absDivisor {
		absDividen = absDividen - absDivisor
		re++
	}

	if (mDd == true && mDr != true) || (mDd != true && mDr == true) {
		return -1 * re
	} else {
		return re
	}
}

func abs(a int) (int, bool) {
	if a < 0 {
		return -1 * a, true
	}

	return a, false
}
*/

func main() {
	d := 7
	r := -3
	fmt.Println(divide(d, r))
}
