/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 *
 * https://leetcode.com/problems/set-mismatch/description/
 *
 * algorithms
 * Easy (41.06%)
 * Likes:    460
 * Dislikes: 252
 * Total Accepted:    54K
 * Total Submissions: 131.4K
 * Testcase Example:  '[1,2,2,4]'
 *
 * 
 * The set S originally contains numbers from 1 to n. But unfortunately, due to
 * the data error, one of the numbers in the set got duplicated to another
 * number in the set, which results in repetition of one number and loss of
 * another number. 
 * 
 * 
 * 
 * Given an array nums representing the data status of this set after the
 * error. Your task is to firstly find the number occurs twice and then find
 * the number that is missing. Return them in the form of an array.
 * 
 * 
 * 
 * Example 1:
 * 
 * Input: nums = [1,2,2,4]
 * Output: [2,3]
 * 
 * 
 * 
 * Note:
 * 
 * The given array size will in the range [2, 10000].
 * The given array's numbers won't have any order.
 * 
 * 
 */


 /*
func findErrorNums(nums []int) []int {
	res := []int{}
	temp := make([]int, len(nums))

	for i:=0; i<len(nums); i++ {
		temp[i] = abs(i+1 - nums[i]) 
	}

	fmt.Println(temp)

	if len(temp) < 3 {
		if temp[0] == 0 {
			res = append(res, 1,2)
		}else if temp[0] == 1{
			res = append(res, 2,1)
		}
		
	}else {
		for i:=0; i<len(temp); i++ {
			if temp[i] ==  1{
				res = append(res, i, i+1)
			}
		}
	}

	fmt.Println(res)

	return res
}
*/

func findErrorNums(nums []int) []int {
	dup := 0
	for i:=0; i<len(nums); i++ {
		n := abs(nums[i])

		if nums[n-1] < 0 {
			dup = n
		}else{
			nums[n-1] = -nums[n-1]
		}
	}

	fmt.Println(nums, dup)

	mis := 0
	for i, v := range nums {
		if v > 0 {
			mis = i + 1
			break
		}
	}

	return []int{dup, mis}
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

