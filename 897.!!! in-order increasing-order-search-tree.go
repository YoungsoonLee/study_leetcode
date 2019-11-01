/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
 *
 * https://leetcode.com/problems/increasing-order-search-tree/description/
 *
 * algorithms
 * Easy (65.80%)
 * Likes:    375
 * Dislikes: 372
 * Total Accepted:    44.4K
 * Total Submissions: 67.1K
 * Testcase Example:  '[5,3,6,2,4,null,8,1,null,null,null,7,9]'
 *
 * Given a binary search tree, rearrange the tree in in-order so that the
 * leftmost node in the tree is now the root of the tree, and every node has no
 * left child and only 1 right child.
 *
 *
 * Example 1:
 * Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]
 *
 * ⁠      5
 * ⁠     / \
 * ⁠   3    6
 * ⁠  / \    \
 * ⁠ 2   4    8
 * /        / \
 * 1        7   9
 *
 * Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
 *
 * ⁠1
 * \
 * 2
 * \
 * 3
 * \
 * 4
 * \
 * 5
 * \
 * 6
 * \
 * 7
 * \
 * 8
 * \
 * ⁠                9
 *
 * Note:
 *
 *
 * The number of nodes in the given tree will be between 1 and 100.
 * Each node will have a unique integer value from 0 to 1000.
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var head = &TreeNode{}
	tail := head
	var rec func(root *TreeNode)
	rec = func(root *TreeNode) {
		if root == nil {
			return
		}
		rec(root.Left)
		root.Left = nil               // 切断 root 与其 Left 的连接，避免形成环
		tail.Right, tail = root, root // 把 root 接上 tail，并保持 tail 指向尾部
		rec(root.Right)
	}
	rec(root)
	return head.Right
}

/*
func rec2(root *TreeNode, tail *TreeNode) {
	if root == nil {
		return nil
	}

	rec2(root.Left, tail)

	root.Left = nil
	tail.Right, tail = root, root

	rec2(root.Right, tail)
}
*/
// @lc code=end

