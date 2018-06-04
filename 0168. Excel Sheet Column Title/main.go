package main

import "fmt"

func convertToTitle(n int) string {
	res := ""

	//fmt.Println(byte(n % 26))

	for n > 0 {
		n--
		fmt.Println(n, string(byte(n%26)+'A')) // 26은 캐릭터 알파벳 수
		res = string(byte(n%26)+'A') + res
		n /= 26
	}

	fmt.Println(res)
	return res
}

func main() {
	a := 15
	convertToTitle(a)
}
