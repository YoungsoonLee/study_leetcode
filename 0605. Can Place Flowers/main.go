package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 && // 检查 i 的值
			((i+1 < l && flowerbed[i+1] == 0) || i+1 >= l) && // 检查 i+1 的值
			((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 < 0) { // 检查 i-1 的值
			flowerbed[i] = 1
			n--
			if n <= 0 {
				return true
			}
		}
	}

	return n <= 0
}

func canPlaceFlowers_my(flowerbed []int, n int) bool {
	// init
	p, c, nn := flowerbed[0], flowerbed[0], flowerbed[0]
	//count := 0

	// ??? why?
	if p == 0 && c == 0 && nn == 0 {
		flowerbed[0] = 1
		n-- // !!!
	}

	for i := 1; i < len(flowerbed)-2; i++ {
		p = flowerbed[i-1]
		c = flowerbed[i]
		nn = flowerbed[i+1]

		if p == 0 && c == 0 && nn == 0 {
			flowerbed[i] = 1
			n--
		}
	}

	fmt.Println(flowerbed)
	fmt.Println(n)

	if n <= 0 {
		return true
	}

	return n <= 0

}

func main() {
	f := []int{1, 0, 0, 0, 1}
	n := 1

	fmt.Println(canPlaceFlowers_my(f, n))

	f = []int{1, 0, 0, 0, 1}
	n = 2

	fmt.Println(canPlaceFlowers_my(f, n))

	f = []int{1, 0, 0, 0, 0, 1}
	n = 2

	fmt.Println(canPlaceFlowers_my(f, n))

}
