package main

import "fmt"

func findNthDigit(n int) int {
	// step 1 : NthDigit 숫자가있는 숫자의 자리 찾기
	// 카운트 자리수 자리수
	count, digits := 9, 1

	// num은 가장 작은 자릿수입니다.
	num := 1
	fmt.Println(n, count, digits, num)
	fmt.Println("n-count * digits : ", n-count*digits)
	for n-count*digits > 0 {
		n -= count * digits
		count *= 10
		digits++
		num *= 10
		fmt.Println(n, count, digits, num, n-count*digits)
	}

	// step 2 : NthDigit으로 숫자 찾기
	// index NthDigit는 대상 번호의 인덱스 번호입니다.
	index := n % digits
	fmt.Println("index: ", index)
	if index == 0 {
		index = digits
	}

	fmt.Println(n, index, digits, num)

	// NthDigit으로 num을 숫자로 만듭니다.
	num += n / digits
	if index == digits {
		num--
	}

	// step 3 : NthDigit 찾기
	for i := index; i < digits; i++ {
		num /= 10
	}

	fmt.Println(num, num%10)
	return num % 10
}

func main() {
	a := 15
	findNthDigit(a)
}
