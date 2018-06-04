package main

import (
	"fmt"
	"strings"
)

func countPrimeSetBits(L int, R int) int {
	primes := [...]int{2: 1, 3: 1, 5: 1, 7: 1, 11: 1, 13: 1, 17: 1, 19: 1} // !!!!!
	fmt.Println(primes)
	p2 := [...]int{1, 2, 3}
	fmt.Println(p2)

	fmt.Println(6 >> 1)
	fmt.Println(6 & 1)
	fmt.Println(3 >> 1)
	fmt.Println(3 & 1)

	res := 0
	for i := L; i <= R; i++ {
		bits := 0
		for n := i; n > 0; n >>= 1 {
			//sfmt.Println(n, bits)
			bits += n & 1
		}
		res += primes[bits]
	}
	fmt.Println(res)
	return res
}

func countPrimeSetBits_my(L int, R int) int {
	count := 0

	for i := L; i <= R; i++ {
		//fmt.Printf("%b\n", i)
		s := fmt.Sprintf("%b\n", i)
		bc := strings.Count(s, "1")

		if isPrime(bc) {
			count++
		}
	}

	fmt.Println(count)
	return count
}

// prime logic
func isPrime(a int) bool {
	fmt.Println(a, a/2)
	if a > 1 && (a/2) > 0 {
		return true
	}
	return false
}

func main() {
	L := 6
	R := 10
	countPrimeSetBits(L, R)
}
