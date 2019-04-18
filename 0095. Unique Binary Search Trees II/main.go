package main

import "fmt"

/**
 * Definition for a binary tree node.
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode

	if n <= 0 {
		return res
	}

	// 소유 조건 조건에 부합하지 않습니다.
	in := make([]int, n)
	for i := range in {
		in[i] = i + 1
	}

	fmt.Println("in: ", in)

	// 사용 inorder 생성 가능 사전 주문
	pres := getPres(in)

	// 선주문과 inorder가있는 이진 트리 생성
	for _, pre := range pres {
		temp := preIn2Tree(pre, in)
		res = append(res, temp)
	}

	return res
}

func preIn2Tree(pre, in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = preIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = preIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}

func getPres(in []int) [][]int {
	size := len(in)
	if size <= 1 {
		return [][]int{in}
	}

	if size == 2 {
		return [][]int{
			[]int{in[1], in[0]},
			[]int{in[0], in[1]},
		}
	}

	res := [][]int{}
	for i := 0; i < size; i++ {
		//와 함께 [i]를 루트로 사용
		// [i]의 왼쪽에있는 선주문을 얻습니다.
		ls := getPres(in[:i])

		// [i]의 오른쪽에있는 선주문을 얻습니다.
		rs := getPres(in[i+1:])

		for _, l := range ls {
			for _, r := range rs {
				temp := make([]int, 1, size)
				// in [i]는 루트이므로 0에 있어야합니다.
				temp[0] = in[i]
				temp = append(temp, l...)
				temp = append(temp, r...)

				// 요약 결과
				res = append(res, temp)
			}
		}
	}

	return res
}

func main() {

}
