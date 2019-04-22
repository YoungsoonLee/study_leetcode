package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 이미 가정
// BST에 중복 값이 없습니다.
// root! = nil
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	var dfs func(*TreeNode)

	// dfs는 중간 순회입니다.
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}

		if prev != nil && prev.Val > root.Val {
			// If you want to adjust 6 and 3 of the misplaced bits in [1, 2, 6, 4, 5, 3, 7]
			// In fact, the larger value in [6, 4] is exchanged with the smaller value in [5, 3]. At this time, there are two sets of misordering.
			// But there may be
			// 3 and 2 of the misplaced bits in [1, 3, 2], only one set of unordered [3, 2]
			// The following two if statements are compatible with the above two cases
			if first == nil {
				first = prev
			}

			if first != nil {
				// When there is a second set of out-of-order
				// the value of second will be modified
				second = root
			}
		}

		prev = root

		if root.Right != nil {
			dfs(root.Right)
		}

	}

	dfs(root)

	// 주제는 교환되는 노드가 있음을 보장합니다.
	// 첫 번째와 두 번째가 nil인지 검사 할 필요가 없다.
	first.Val, second.Val = second.Val, first.Val
}

func main() {

}
