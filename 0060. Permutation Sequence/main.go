package main

import (
	"fmt"
)

// 퍼뮤테이션은 재귀 !!!
func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	//결과를 저장하는 데 사용되는 문자
	res := make([]byte, n)

	// 크롤링 할 번호 저장
	rec := make([]byte, n)
	for i := 0; i < n; i++ {
		rec[i] = byte(i) + '1'
	}

	fmt.Println("rec: ", string(rec))

	// 시퀀스 번호가 1부터 시작하기 때문입니다.
	// k는 0부터 시작하려면 1을 빼야합니다.
	k--
	base := 1
	for i := 2; i < n; i++ {
		base *= i
	}

	fmt.Println("base: ", base)

	for i := 0; i < n-1; i++ {
		idx := k / base
		res[i] = rec[idx]

		fmt.Println(idx, k, base, string(res), string(rec))

		// rec에서 rec [idx]를 제거한다.
		rec = append(rec[:idx], rec[idx+1:]...)
		k %= base
		base /= (n - i - 1)
	}
	// 마지막 번호를 잊지 마라.
	res[n-1] = rec[0]
	fmt.Println("res: ", string(res))
	return string(res)
}

func main() {
	n := 3
	k := 2
	getPermutation(n, k)
}
