package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeWithDepth struct {
	node  *TreeNode
	depth int
}

// BST
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := make([]NodeWithDepth, 0, 1024)
	q = append(q, NodeWithDepth{root, 1})
	maxd := 0

	for len(q) > 0 {
		nd := q[0]
		if nd.depth > maxd {
			maxd = nd.depth
		}
		if nd.node.Left != nil {
			q = append(q, NodeWithDepth{nd.node.Left, nd.depth + 1})
		}
		if nd.node.Right != nil {
			q = append(q, NodeWithDepth{nd.node.Right, nd.depth + 1})
		}
		q = q[1:]
	}
	return maxd
}

// DFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxD = 1
	if root.Left != nil {
		dfs(root.Left, 2)
	}
	if root.Right != nil {
		dfs(root.Right, 2)
	}
	return maxD
}

var maxD int

func dfs(node *TreeNode, depth int) {
	if depth > maxD {
		maxD = depth
	}

	if node.Left != nil {
		dfs(node.Left, depth+1)
	}

	if node.Right != nil {
		dfs(node.Right, depth+1)
	}
}

// DFS 2
func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxv := 0

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		maxv = max(maxv, depth)

		if root.Left != nil {
			dfs(root.Left, depth+1)
		}
		if root.Right != nil {
			dfs(root.Right, depth+1)
		}
	}
	dfs(root, 1)
	return maxv
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func maxDepth(root *TreeNode) int {
	//max := 0
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
*/

/*
func arrToTree(arr []int) *TreeNode {
	t := &TreeNode{Val: arr[0], Left: nil, Right: nil}

	for _, i := range arr[1:] {
		// TODO:
	}

	return t
}
*/

func main() {

}
