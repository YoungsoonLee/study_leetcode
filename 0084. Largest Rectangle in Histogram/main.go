package main

import "fmt"

func largestRectangleArea(heights []int) int {
	// 원래는 높이는 음수가 아닙니다.
	// -2와 -1을 추가 한 후 for 루프에서 시작하는 계산을 단순화합니다.
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)
	size := len(heights)

	// 높이 요소의 인덱스가 스택에 저장됩니다.
	// 스택의 인덱스 번호에 해당하는 높이의 값이 증가합니다.
	// 예 :
	// stack = [] int {1,3,5}, 그런 다음
	// 높이 [1] <높이 [3] <높이 [5]
	stack := make([]int, 1, size)
	// 높이 [0]의 인덱스 번호를 스택에 넣습니다.
	// 1은 1로 시작합니다.
	end := 1

	//fmt.Println(stack)

	res := 0
	for end < size {
		// 끝점을 새 높이로 지정하고 스택에 끝을 넣고 다음을 가리킴
		if heights[stack[len(stack)-1]] < heights[end] {
			stack = append(stack, end)
			end++
			continue
		}

		begin := stack[len(stack)-2]
		index := stack[len(stack)-1]
		height := heights[index]

		fmt.Println(begin, index, height)

		// area는 높이 사이의 가장 큰 사각형의 영역입니다 (begin : end). 왜냐하면
		// 높이의 누구 (시작 : 색인)> 높이> 높이의 사람 (색인 : 끝)
		area := (end - begin - 1) * height
		res = max(res, area)
		// h의 인덱스 번호는 더 이상 스택에 남겨 둘 필요가 없습니다.
		stack = stack[:len(stack)-1]
	}

	//fmt.Println(stack, res)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := []int{2, 1, 5, 6, 2, 3}
	largestRectangleArea(n)
}
