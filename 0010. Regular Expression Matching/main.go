package main

import "fmt"

// The following assumptions exist in the program
// "*" does not appear in the first position of p
// "**" will not appear, but ".*."" , ".*.." , ".*.*" will appear

/*
func isMatch_recur(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := !isEmpty(s) && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return (isMatch_recur(s, p[2:])) || (firstMatch && isMatch_recur(s[1:], p))
	} else {
		return firstMatch && isMatch_recur(string(s[1]), string(p[1]))
	}
}

func isEmpty(s string) bool {
	return len(s) == 0
}
*/

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1) // row
	for i := range dp {
		dp[i] = make([]bool, len(s)+1) // column
	}

	dp[0][0] = true
	//fmt.Println(len(dp))

	for i := 2; i < len(dp); i += 2 {
		if p[i-1] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}

	fmt.Println("1: ", dp)

	for i := 1; i < len(dp); i++ {
		if i < len(p) && p[i] == '*' {
			continue
		}

		for j := 1; j < len(dp[0]); j++ {
			fmt.Println(string(p[i-1]))
			if p[i-1] == '*' {
				if p[i-2] == '.' {
					dp[i][j] = dp[i-2][j-1] || dp[i][j-1] || dp[i-2][j]
				} else {
					dp[i][j] = (dp[i-2][j]) || (p[i-2] == s[j-1] && (dp[i-2][j-1] || dp[i][j-1]))
				}
			} else if p[i-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] && p[i-1] == s[j-1]
			}
		}
	}

	return dp[len(p)][len(s)]
}

func main() {
	s := "aa"
	p := "a"
	fmt.Println(isMatch(s, p))
}
