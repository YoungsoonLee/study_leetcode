import "fmt"

/*
 * @lc app=leetcode id=922 lang=golang
 *
 * [922] Sort Array By Parity II
 *
 * https://leetcode.com/problems/sort-array-by-parity-ii/description/
 *
 * algorithms
 * Easy (67.22%)
 * Likes:    381
 * Dislikes: 37
 * Total Accepted:    65K
 * Total Submissions: 96.6K
 * Testcase Example:  '[4,2,5,7]'
 *
 * Given an array A of non-negative integers, half of the integers in A are
 * odd, and half of the integers are even.
 *
 * Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is
 * even, i is even.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [4,2,5,7]
 * Output: [4,5,2,7]
 * Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been
 * accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 2 <= A.length <= 20000
 * A.length % 2 == 0
 * 0 <= A[i] <= 1000
 *
 *
 *
 *
 *
 */

// @lc code=start
func sortArrayByParityII(A []int) []int {
	size := len(A)
	res := make([]int, size)
	even, odd := 0, 1

	for _, a := range A {
		if a%2 == 0 {
			res[even] = a
			even += 2
		} else {
			res[odd] = a
			odd += 2
		}
	}

	fmt.Println(res)
	return res

	/* mine
	odd := make([]int, 0, len(A)/2)
	even := make([]int, 0, len(A)/2)

	for i := 0; i < len(A); i++ {
		if i%2 != 0 {
			if isEven(A[i]) {
				even = append(even, A[i])
				A[i] = -1
			}
		} else if i%2 == 0 {
			if !isEven(A[i]) {
				odd = append(odd, A[i])
				A[i] = -1
			}
		}
	}

	fmt.Println(odd)
	fmt.Println(even)

	for i := 0; i < len(A); i++ {
		if i%2 != 0 && A[i] == -1 {
			A[i] = odd[0]
			odd = odd[1:]
		} else if i%2 == 0 && A[i] == -1 {
			A[i] = even[0]
			even = even[1:]
		}
	}
	return A
	*/
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

// @lc code=end

