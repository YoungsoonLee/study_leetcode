/*
 * @lc app=leetcode id=1042 lang=golang
 *
 * [1042] Flower Planting With No Adjacent
 *
 * https://leetcode.com/problems/flower-planting-with-no-adjacent/description/
 *
 * algorithms
 * Easy (47.87%)
 * Likes:    184
 * Dislikes: 225
 * Total Accepted:    16.4K
 * Total Submissions: 34.2K
 * Testcase Example:  '3\n[[1,2],[2,3],[3,1]]'
 *
 * You have N gardens, labelled 1 to N.  In each garden, you want to plant one
 * of 4 types of flowers.
 * 
 * paths[i] = [x, y] describes the existence of a bidirectional path from
 * garden x to garden y.
 * 
 * Also, there is no garden that has more than 3 paths coming into or leaving
 * it.
 * 
 * Your task is to choose a flower type for each garden such that, for any two
 * gardens connected by a path, they have different types of flowers.
 * 
 * Return any such a choice as an array answer, where answer[i] is the type of
 * flower planted in the (i+1)-th garden.  The flower types are denoted 1, 2,
 * 3, or 4.  It is guaranteed an answer exists.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: N = 3, paths = [[1,2],[2,3],[3,1]]
 * Output: [1,2,3]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: N = 4, paths = [[1,2],[3,4]]
 * Output: [1,2,1,2]
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: N = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
 * Output: [1,2,3,4]
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= N <= 10000
 * 0 <= paths.size <= 20000
 * No garden has 4 or more paths coming into or leaving it.
 * It is guaranteed an answer exists.
 * 
 * 
 * 
 * 
 */

// @lc code=start
func gardenNoAdj(N int, paths [][]int) []int {
	connects := make([][]int, N)
	
	// ??? 뭘하는 로직일까...
	for _, p := range paths {
		i,j := p[0]-1, p[1]-1 //i,j = x-1, y-1
		connects[i] = append(connects[i], j)
		connects[j] = append(connects[j], i)
		fmt.Println(connects)
	}

	fmt.Println(connects)

	res := make([]int, N)
	
	for i:=0; i<N; i++ {
		isUsed := [5]bool{}
		for _,j := range connects[i] {
			isUsed[res[j]] = true
		}
		for color := 1; color <= 4; color++ {
			if !isUsed[color] {
				res[i] = color
				break
			}
		}
	}

	return res


}
// @lc code=end

