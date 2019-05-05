package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil {
		return nil
	}

	t := &TreeNode{Val: postorder[len(postorder)-1]}

	if len(inorder) == 1 {
		return t
	}

	idx := indexOf(t.Val, inorder)

	t.Left = buildTree(inorder[:idx], postorder[:idx])
	t.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])

	return t

}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}

func main() {

}
