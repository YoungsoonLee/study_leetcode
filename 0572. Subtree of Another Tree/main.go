package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func isSubtree(s *TreeNode, t *TreeNode) bool {
	return DFSFunc(s, t, compareTree)
}

func DFSFunc(s, t *TreeNode, f func(*TreeNode, *TreeNode) bool) bool {
	if s == nil {
		if t == nil {
			return true
		}
		return false
	}
	if f(s, t) == true {
		return true
	}
	if DFSFunc(s.Left, t, f) == true {
		return true
	}

	return DFSFunc(s.Right, t, f)
}

func compareTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		if t2 == nil {
			return true
		}
		return false
	} else if t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}
	if compareTree(t1.Left, t2.Left) {
		return false
	}
	return compareTree(t1.Right, t2.Right)

}
*/

/*
func isSubtree(s *TreeNode, t *TreeNode) bool {
	return DFSFunc(s, t, compareTree)
}

func DFSFunc(s, t *TreeNode, f func(*TreeNode, *TreeNode) bool) bool {
	if s == nil {
		if t == nil {
			return true
		}
		return false
	}

	if f(s, t) == true {
		return true
	}

	if DFSFunc(s.Left, t, f) == true {
		return true
	}
	return DFSFunc(s.Right, t, f)
}

func compareTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		if t2 == nil {
			return true
		}
		return false
	} else if t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}

	if !compareTree(t1.Left, t2.Left) {
		return false
	}
	return compareTree(t1.Right, t2.Right)
}
*/

/*
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSame(s, t) {
		return true
	}

	if s == nil {
		return false
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t) // !!!!!
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}

	// 이미 s는 nil이 아니므로...
	if t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right) // !!!
}
*/

/*
func isSubtree(s *TreeNode, t *TreeNode) bool {
	s = pivotStree(s, t)

	if s.Val != t.Val {
		return false
	}

	return compareTree(s.Left, t) && compareTree(s.Right, t)
}

func compareTree(s *TreeNode, t *TreeNode) bool {
	if s.Val != t.Val {
		return false
	}

	return true
}

func pivotStree(s *TreeNode, t *TreeNode) *TreeNode {
	if s.Val != t.Val {
		if s.Left != nil {
			pivotStree(s.Left, t)
		}
		if s.Right != nil {
			pivotStree(s.Right, t)
		}
	}
	return s
}
*/

func main() {}
