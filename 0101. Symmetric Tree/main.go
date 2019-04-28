package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	q := make([]*TreeNode, 0)

	q = append(q, root)
	q = append(q, root)

	for len(q) != 0 {
		var t1 *TreeNode
		t1, q = q[0], q[1:] //pop

		var t2 *TreeNode
		t2, q = q[0], q[1:] //pop

		if t1 == nil && t2 == nil {
			continue
		}

		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}

		q = append(q, t1.Left)
		q = append(q, t2.Right)

		q = append(q, t1.Right)
		q = append(q, t2.Left)

	}

	return false
}

func main() {

}
