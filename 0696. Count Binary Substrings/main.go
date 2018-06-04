package main

func countBinarySubstrings(s string) int {
	count, countZero, countOne := 0, 0, 0
	prev := rune(s[0])

	for _, r := range s {
		if prev == r {
			if r == '0' {
				countZero++
			} else {
				countOne++
			}
		} else {
			// 더 적은 수는 질문과 일치하는 부분 문자열의 수를 결정합니다.
			// 예를 들어, "00011"은 "0011", "01"의 하위 문자열이며 첫 번째 "0"은 쓸모가 없습니다
			count += min(countZero, countOne)
			if r == '0' {
				countZero = 1
			} else {
				countOne = 1
			}
		}
		prev = r
	}

	return count + min(countZero, countOne)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {}
