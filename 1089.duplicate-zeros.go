import "fmt"

/*
 * @lc app=leetcode id=1089 lang=golang
 *
 * [1089] Duplicate Zeros
 *
 * https://leetcode.com/problems/duplicate-zeros/description/
 *
 * algorithms
 * Easy (58.48%)
 * Likes:    273
 * Dislikes: 142
 * Total Accepted:    28.7K
 * Total Submissions: 48.8K
 * Testcase Example:  '[1,0,2,3,0,4,5,0]'
 *
 * Given a fixed length array arr of integers, duplicate each occurrence of
 * zero, shifting the remaining elements to the right.
 *
 * Note that elements beyond the length of the original array are not written.
 *
 * Do the above modifications to the input array in place, do not return
 * anything from your function.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,0,2,3,0,4,5,0]
 * Output: null
 * Explanation: After calling your function, the input array is modified to:
 * [1,0,0,2,3,0,0,4]
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3]
 * Output: null
 * Explanation: After calling your function, the input array is modified to:
 * [1,2,3]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= arr.length <= 10000
 * 0 <= arr[i] <= 9
 *
 */

// @lc code=start
func duplicateZeros(arr []int) {
	n := len(arr)

	// count
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			count++
		}
	}

	// copy A[i] to A[j]
	copy := func(i, j int) {
		if j < n {
			arr[j] = arr[i]
		}
	}

	//fmt.Println(count)
	i, j := n-1, n+count-1
	//fmt.Println(i, j)

	for i < j {
		copy(i, j)
		fmt.Println(arr)
		if arr[i] == 0 {
			j--
			copy(i, j)
		}
		i--
		j--
	}

}

// @lc code=end

