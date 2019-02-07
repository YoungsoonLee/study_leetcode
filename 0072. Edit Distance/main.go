package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	// dp init
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	fmt.Println("init dp: ", dp)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// dp [i] [j]는 [: i] → [: j]까지 필요한 최소 단계를 나타냅니다.
			// 제목에 주어진 세 가지 작동 방법에 따라 별도로 토론하십시오.
			// 1. 	[: i-1]에서 얻을 마지막 글자의 첫 번째 * 삭제 *.
			// 		Then [: i-1] → to [: j]이 메서드에 필요한 단계는 다음과 같습니다.
			// 		1 + dp [i-1] [j]
			// 2. [: i] → [: j-1]로,
			// 		그런 다음 *의 [j-1]을 to의 끝에 추가하면이 메소드에 필요한 단계는 다음과 같습니다.
			// 		1 + dp [i] [j-1]
			// 3. [i-1] → [j-1]
			// 		3.1 [i-1] =에서 [i-1]
			// 			*를 바꿀 필요가 없습니다. *,
			// 			총 단계는 dp [i-1] [j-1]
			// 		3.2에서 [i-1]! = ~ [i-1]
			// 			[replace] 작업을 [i-1]에서 [j-1]로 바꿉니다.
			// 			총 단계는 1 + dp [i-1] [j-1]
			dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])

			replace := 1
			if word1[i-1] == word2[j-1] {
				replace = 0
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+replace)
		}
	}

	fmt.Println("dp: ", dp)
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	w1 := "intention"
	w2 := "execution"
	minDistance(w1, w2)
}

/*
func minDistance(word1 string, word2 string) int {
	w1l := len(word1)
	w2l := len(word2)

	hash := make(map[byte]byte, 0) // TODO: reduce capacity.

	if w1l == w2l {
		for i := 0; i < w1l; i++ {
			if _, ok := hash[word1[i]]; ok {
				//do nothing
			} else {
				if word1[i] != word2[i] || word1[i] == word2[i] {
					hash[word1[i]] = word2[i]
				}
			}
		}
	}

	if w1l > w2l {
		for i := 0; i < w1l; i++ {
			if _, ok := hash[word1[i]]; ok {
				//do nothing
			} else {
				if i < w2l {
					if word1[i] != word2[i] || word1[i] == word2[i] {
						hash[word1[i]] = word2[i]
					}
				}

			}
		}
	}

	if w1l < w2l {
		for i := 0; i < w2l; i++ {
			if _, ok := hash[word1[i]]; ok {
				//do nothing
			} else {
				if i < w1l {
					if word1[i] != word2[i] || word1[i] == word2[i] {
						hash[word1[i]] = word2[i]
					}
				}

			}
		}
	}

	fmt.Println(hash, len(hash))
	return len(hash)
}

func main() {
	w1 := "intention"
	w2 := "execution"
	minDistance(w1, w2)
}
*/
