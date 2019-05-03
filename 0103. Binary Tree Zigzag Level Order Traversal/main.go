package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level > len(res) {
			res = append(res, []int{})
		}

		// 102 문제와 비교하면 레벨의 패리티가 다르므로 추가가 달라집니다.
		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			temp := make([]int, len(res[level])+1)
			temp[0] = root.Val
			copy(temp[1:], res[level])

			fmt.Println(temp)
			fmt.Println(res)

			res[level] = temp
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)

	}

	dfs(root, 0)
	return res
}

func main() {

}
