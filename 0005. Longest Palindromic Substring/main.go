package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// 가장 긴 회유 문장의 첫 번째 문자 색인과 가장 긴 회문의 길이
	begin, maxLen := 0, 1

	// for 루프에서
	// b는 회문에서 ** 문자의 색인을 나타내며,
	// e는 회문의 ** 꼬리 ** 문자 색인 번호를 나타내며,
	// i는 회문의 중간 부분의 첫 번째 문자의 색인 번호를 나타냅니다.
	// 각 for 루프에서
	// 먼저 i로 시작하고, '중간 중간 부분'의 모든 문자의 동일한 특성을 사용하고, b, e가 각각 '중간 중간 섹션'의 시작과 끝을 가리 키도록합니다.
	// 중간 중간에서 중간으로 확장하여 b, e가이 중 중간 부분을 중심으로 가장 긴 회궁의 끝을 가리 키도록합니다.
	for i := 0; i < len(s); { // s [i]가 '중간 중간'의 첫 번째 문자로 가장 긴 회중을 찾습니다.
		if len(s)-i <= maxLen/2 {
			// 왜냐하면 i는 회중 색인의 첫 번째 문자의 색인 번호이므로 "중간 중간"
			//이 때 발견 될 수있는 가장 긴 회상이 l이라면 l <= (len (s) -i) * 2 - 1
			// len (s) -i <= maxLen / 2 인 경우
			// 그런 다음, l <= maxLen - 1
			// 그런 다음, l <maxLen
			// 다시 찾지 않아도됩니다.
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// 루프가 끝나면 s [b : e + 1]은 같은 문자열입니다.
		}

		fmt.Println("first: ", b, e)

		// 다음 회문의 '중간 중간'의 첫 번째 문자는 s [e + 1]
		// 다음 사이클을 준비한다.
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// 루프가 끝나면 s [b : e + 1]은 이번에 발견 할 수있는 가장 긴 회문입니다.
		}

		fmt.Println("second: ", b, e)
		newLen := e + 1 - b

		fmt.Println("third: ", newLen, maxLen, b)

		// 혁신 레코드가 있으면 레코드 업데이트
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	fmt.Println("forth: ", begin, maxLen)
	return s[begin : begin+maxLen]
}

func main() {
	s := "cbbd"
	longestPalindrome(s)
}
