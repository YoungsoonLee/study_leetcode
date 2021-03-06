package main

import "fmt"

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	ret := []string{""}

	for i := 0; i < len(digits); i++ {
		temp := []string{}
		// ret의 각 요소에 새 문자를 추가합니다.
		for j := 0; j < len(ret); j++ {
			// ret [j] 끝에 여러 번 [i] 자에 해당하는 글자를 추가합니다.
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}

		fmt.Println(ret, temp)
		ret = temp

	}

	return ret

}

func main() {
	d := "234"
	fmt.Println(letterCombinations(d))
}
