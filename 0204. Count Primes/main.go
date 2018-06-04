package main

import "fmt"

func countPrimes(n int) int {
	count := 0

	for i := 1; i < n-1; i++ {
		if n%i != 0 {
			count++
		}
	}
	fmt.Println(count)
	return count
}

func countPrimes2(n int) int {
	if n < 3 {
		return 1
	}

	isComposite := make([]bool, n)
	// n보다 작은 숫자들 중 절반은 짝수이고 나머지 2 개는 소수가 아닙니다.
	// 따라서 상한값은 n / 2입니다.
	count := n / 2

	for i := 3; i*i < n; i += 2 {
		fmt.Println("i: ", i)
		if isComposite[i] {
			continue
		}

		// j는 i의 배수이므로 j는 확실히 소수가 아닙니다.
		// 모두 j로 표시

		for j := i * i; j < n; j += 2 * i {
			fmt.Println("j: ", j)
			if !isComposite[j] {
				// j는 다른 태그로 태그 지정되지 않았습니다.
				// 개수를 수정하십시오.
				count--
				// mark j
				isComposite[j] = true
			}
		}
	}

	fmt.Println(isComposite)
	fmt.Println(count)

	return count
}

func main() {
	countPrimes2(25)
}
