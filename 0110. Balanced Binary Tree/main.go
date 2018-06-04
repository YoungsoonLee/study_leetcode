/**
 * height-balanced binary
 * a binary tree in which the depth of the two subtrees of every node never differ by more than 1
 * 모든 노드의 두 하위 트리의 깊이가 1을 넘지 않는 이진 트리
 *
 **/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// my solution - fail
/*
func isBalanced(root *TreeNode) bool {
	leftCount := 0
	rightCount := 0

	var leftT *TreeNode
	var righT *TreeNode

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil {
		leftCount++
		leftT = root.Left

		for leftT.Left != nil {
			leftCount++
			leftT = leftT.Left
		}
	}

	if root.Right != nil {
		rightCount++
		righT = root.Right

		for righT.Right != nil {
			rightCount++
			righT = righT.Right
		}
	}
	fmt.Println(leftCount, rightCount)
	return int(math.Abs(float64(leftCount)-float64(rightCount))) < 2
}
*/

func isBalanced(root *TreeNode) bool {
	_, isBalanced := recur(root)
	return isBalanced
}

func recur(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := recur(root.Left)
	rightDepth, rightIsBalanced := recur(root.Right)

	if leftIsBalanced && rightIsBalanced && -1 <= leftDepth-rightDepth && leftDepth-rightDepth <= 1 {
		//fmt.Println(leftDepth, rightDepth)
		return max(leftDepth, rightDepth) + 1, true
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//t2Left := &TreeNode{Val: 15, Left: nil, Right: nil}
	//t2Right := &TreeNode{Val: 7, Left: nil, Right: nil}

	t1Left := &TreeNode{Val: 2, Left: nil, Right: nil}
	//t1Right := &TreeNode{Val: 20, Left: t2Left, Right: t2Right}

	t1Root := &TreeNode{Val: 1, Left: t1Left, Right: nil}

	print(isBalanced(t1Root))
}
