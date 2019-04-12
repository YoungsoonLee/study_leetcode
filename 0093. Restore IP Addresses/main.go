package main

import "fmt"

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}

	res := []string{}
	combination := make([]string, 4)

	fmt.Println("combination: ", combination)

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		// fmt.Println(idx, begin)

		if idx == 3 {
			temp := s[begin:]
			if isOK(temp) {
				combination[3] = temp
				res = append(res, IP(combination))
			}
			return
		}

		// 剩余 IP 段，最多需要的宽度
		maxRemain := 3 * (3 - idx)

		for end := begin + 1; end <= n-(3-idx); end++ {
			if end+maxRemain < n {
				// 다음 IP 세그먼트 중 하나 이상이 3자를 초과합니다.
				// 이 IP 세그먼트가 짧다는 것을 설명합니다.
				continue
			}

			if end-begin > 3 {
				//이 IP 세그먼트가 3 자리 이상입니다.
				break
			}

			temp := s[begin:end]
			if isOK(temp) {
				combination[idx] = temp
				dfs(idx+1, end) // !!!!!
			}
		}
	}

	dfs(0, 0)

	return res
}

// IP는 s로 표시된 IP 주소를 반환합니다.
func IP(s []string) string {
	return fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
}

func isOK(s string) bool {
	// 프로그램의 다른 부분에 의해 보증 됨 len (s) <= 3
	// "0"수 있습니다.
	// "01"은 할 수 없습니다.
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	if len(s) < 3 {
		return true
	}

	// len(s) == 3
	switch s[0] {
	case '1':
		return true
	case '2':
		if '0' <= s[1] && s[1] <= 4 {
			return true
		}
		if s[1] == '5' && '0' <= s[2] && s[2] <= '5' {
			return true
		}
	}

	return false
}

func main() {
	//s := "25525511135"
	s := "25251135"
	fmt.Println(restoreIpAddresses(s))
}
