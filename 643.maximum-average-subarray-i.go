/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 *
 * https://leetcode.com/problems/maximum-average-subarray-i/description/
 *
 * algorithms
 * Easy (40.12%)
 * Likes:    507
 * Dislikes: 89
 * Total Accepted:    58.4K
 * Total Submissions: 145.5K
 * Testcase Example:  '[1,12,-5,-6,50,3]\n4'
 *
 * Given an array consisting of n integers, find the contiguous subarray of
 * given length k that has the maximum average value. And you need to output
 * the maximum average value.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,12,-5,-6,50,3], k = 4
 * Output: 12.75
 * Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= k <= n <= 30,000.
 * Elements of the given array will be in the range [-10,000, 10,000].
 * 
 * 
 * 
 * 
 */
func findMaxAverage(nums []int, k int) float64 {
	maxAvg := 0.0
	l := len(nums)
	s := (l - k) + 1
	//c := 1
	sum := 0

	fmt.Println(l, s)
	temp := make([]int, k)

	for i := 0; i < s; i++ {
		//fmt.Println(nums[l-k-i : l-i])
		temp = nums[l-k-i : l-i]
		fmt.Println(temp)

		for j := 0; j < len(temp); j++ {
			sum += temp[j]
		}

		if maxAvg < float64(sum/k) {
			maxAvg = float64(sum) / float64(k)
		}

		fmt.Println("sum: ", sum, float64(sum)/float64(k), maxAvg)
		sum = 0
	}
	//fmt.Println(maxAvg)
	return maxAvg
}

