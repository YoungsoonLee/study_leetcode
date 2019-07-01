package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	// TODO: check denominator zero
	if numerator == 0 {
		return "0"
	}

	if numerator*denominator < 0 {
		// n, d 异号，则结果前面有负号
		return "-" + fractionToDecimal(abs(numerator), abs(denominator))
	}

	numerator, denominator = abs(numerator), abs(denominator)

	if numerator >= denominator {
		fmt.Println("??: ", numerator%denominator)
		// n / d의 소수 부분
		ds := fractionToDecimal(numerator%denominator, denominator)

		fmt.Println("???: ", numerator/denominator, ds)

		return strconv.Itoa(numerator/denominator) + ds[1:]
	}

	//fmt.Println("??: ", numerator%denominator)

	// digits는 n / d 결과를 저장하는 데 사용됩니다.
	digits := make([]byte, 2, 1024)
	digits[0] = '0'
	digits[1] = '.'
	// idx는 숫자로 된 n / d 색인 번호의 결과입니다.
	idx := 2
	rec := make(map[int]int, 1024)
	for {
		if i, ok := rec[numerator]; ok {
			// n n / d가 다음 루프의 시작임을 나타냅니다.
			// 루프 부분의 시작점은 n의 마지막 idx 값입니다.
			return fmt.Sprintf("%s(%s)", string(digits[:i]), string(digits[i:]))
		}

		rec[numerator] = idx
		numerator *= 10
		idx++

		digits = append(digits, byte(numerator/denominator)+'0')
		numerator %= denominator

		if numerator == 0 {
			// 루프가 없습니다.
			return string(digits)
		}
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func main() {
	n := 6
	d := 2
	fractionToDecimal(n, d)
}
