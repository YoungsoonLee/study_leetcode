package main

import "fmt"

func findSubstring(s string, words []string) []int {
	lens := len(s)
	res := make([]int, 0, lens)

	lenws := len(words)
	if lens == 0 || lenws == 0 || lens < lenws*len(words[0]) {
		return res
	}

	lenw := len(words[0])
	// 각 단어의 총 발생 횟수를 단어로 기록합니다.
	record := make(map[string]int, lenws)
	for _, w := range words {
		if len(w) != lenw {
			// 새로 추가 된 예제 2는 단어 길이가 다릅니다.
			// 이것은 가정을 위반한다.
			// 직접 res를 반환합니다.
			return res
		}
		record[w]++
	}
	fmt.Println("record: ", record)
	//각 단어가 단어로 나타날 수있는 횟수
	remain := make(map[string]int, lenws)
	// 요구 사항을 만족하는 단어가 연속적으로 나오는 횟수를 계산합니다.
	count := 1 // count가 아닌 경우 count의 초기 값은 0이 될 수 있습니다.
	left, right := 0, 0

	reset := func() {
		if count == 0 {
			// 통계 레코드가 수정되지 않았으므로 다시 설정할 필요가 없습니다.
			//이 사전 판단 때문에 재설정의 첫 번째 실행이 필요합니다.
			// count의 값은 0이 될 수 없다.
			// count의 초기 값은 0이 될 수 없다.
			return
		}
		for k, v := range record {
			remain[k] = v
		}
		count = 0
		fmt.Println("reset remain: ", remain)
	}

	// moveLeft
	moveLeft := func() {
		remain[s[left:left+lenw]]++
		count--
		left += lenw
	}

	// 놓치지 않기 위해 {0,1, ..., lenw-1}에서 검사를 시작해야합니다.
	for i := 0; i < lenw; i++ {
		left, right = i, i
		reset()
		// s의 길이 [left :]> = 단어의 모든 단어의 문자열 길이
		// 찾을 문자열이있을 수 있습니다. [left :]
		for lens-left >= lenws*lenw {
			word := s[right : right+lenw]
			remainTimes, ok := remain[word]

			switch {
			case !ok:
				// word 不在 words 中
				// 从 right+lenw 处，作为新窗口，重新开始统计
				left, right = right+lenw, right+lenw
				reset()
			case remainTimes == 0:
				// word 的出现次数上次就用完了
				// 说明 word 在 s[left:right] 中出现次数过多
				moveLeft()
				// 这个case会连续出现
				// 直到 s[left:right] 中的统计结果是 remain[word] == 1
				// 这个过程中，right 一直不变
			default:
				// ok && remainTimes > 0，word 符合出现的条件
				// moveRight
				remain[word]--
				count++
				right += lenw
				// 检查 words 能否排列组合成 s[left:right]
				if count == lenws {
					res = append(res, left)
					// moveLeft 可以避免重复统计 s[left+lenw:right] 中的信息
					moveLeft()
				}
			}
		}

	}
	return res
}

func main() {
	s := "barfoothefoobarman"
	w := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, w))
}
