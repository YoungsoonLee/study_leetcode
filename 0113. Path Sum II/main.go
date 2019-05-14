package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		// 레벨에 따라 경로 업데이트
		// !!!!!
		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}

		sum -= node.Val

		// 잎에 닿는다.
		//이 경로는 요구 사항을 충족합니다.
		if node.Left == nil && node.Right == nil && sum == 0 {
			temp := make([]int, level+1)
			// len (경로)> len (임시) 인 경우 복사가 경로 [len (temp) :]를 복사하지 않습니다.
			copy(temp, path)
			res = append(res, temp)
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}

	dfs(root, 0, sum)

	/*
		var dfs func(*TreeNode, int, *[]int, *[][]int)
		dfs = func(root *TreeNode, sum int, temp *[]int, res *[][]int) {
			if root == nil {
				return
			}

			sum -= root.Val
			*temp = append(*temp, root.Val)

			if sum == 0 {
				*res = append(*res, *temp)
				return
			}

			dfs(root.Left, sum, temp, res)
			dfs(root.Right, sum, temp, res)
		}

		dfs(root, sum, &temp, &res)
	*/
	return res
}

func main() {

}
