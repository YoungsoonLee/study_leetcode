package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)

	fmt.Println("res: ", res)
	fmt.Println("bytes: ", bytes)

	dfs(n, n, 0, bytes, &res)

	fmt.Println("res: ", res)
	return res
}

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	// 모든 심볼이 추가됩니다.
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
	}

	// "("일치하는 것에 대해 걱정하지 마십시오.
	// 직접 추가하려면 왼쪽> 0을 더하십시오.

	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}

	// 추가 할 때 ")"
	// 왼쪽에 필요 <오른쪽,
	// 즉, 적어도 하나의 "("이 "와 일치 할 수 있음)"을 바이트 [: idx]
	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}

func main() {
	n := 3
	generateParenthesis(n)
}
