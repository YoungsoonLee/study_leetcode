import "fmt"

/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
 *
 * https://leetcode.com/problems/leaf-similar-trees/description/
 *
 * algorithms
 * Easy (63.89%)
 * Likes:    473
 * Dislikes: 26
 * Total Accepted:    59.8K
 * Total Submissions: 93.2K
 * Testcase Example:  '[3,5,1,6,2,9,8,null,null,7,4]\n' +
  '[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]'
 *
 * Consider all the leaves of a binary tree.  From left to right order, the
 * values of those leaves form a leaf value sequence.
 *
 *
 *
 * For example, in the given tree above, the leaf value sequence is (6, 7, 4,
 * 9, 8).
 *
 * Two binary trees are considered leaf-similar if their leaf value sequence is
 * the same.
 *
 * Return true if and only if the two given trees with head nodes root1 and
 * root2 are leaf-similar.
 *
 *
 *
 * Note:
 *
 *
 * Both of the given trees will have between 1 and 100 nodes.
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 제목에서 알 수 있듯이 루트에는 노드가 100 개를 넘지 않습니다.
	// 리프 노드는 100을 초과하지 않습니다.
	// a는 항상 s의 기본 배열이며 변경되지 않습니다.
	// 최종 비교는 a1과 a2의 유사점과 차이점을 비교합니다.
	a1 := [100]int{}
	s1 := a1[:0]
	//fmt.Println(a1)
	//fmt.Println(s1)
	search(root1, &s1)

	a2 := [100]int{}
	s2 := a2[:0]
	//fmt.Println(a1)
	//fmt.Println(s1)
	search(root2, &s2)

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(a1)
	fmt.Println(a2)

	return a1 == a2
}

// !!!!
func search(root *TreeNode, sp *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		// leaf를 담는다.
		*sp = append(*sp, root.Val)
		return // !!!
	}

	search(root.Left, sp)
	search(root.Right, sp)
}

// @lc code=end

