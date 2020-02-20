/*
 * @lc app=leetcode id=1103 lang=golang
 *
 * [1103] Distribute Candies to People
 *
 * https://leetcode.com/problems/distribute-candies-to-people/description/
 *
 * algorithms
 * Easy (60.07%)
 * Likes:    161
 * Dislikes: 44
 * Total Accepted:    18.1K
 * Total Submissions: 30.2K
 * Testcase Example:  '7\n4'
 *
 * We distribute some number of candies, to a row of n = num_people people in
 * the following way:
 * 
 * We then give 1 candy to the first person, 2 candies to the second person,
 * and so on until we give n candies to the last person.
 * 
 * Then, we go back to the start of the row, giving n + 1 candies to the first
 * person, n + 2 candies to the second person, and so on until we give 2 * n
 * candies to the last person.
 * 
 * This process repeats (with us giving one more candy each time, and moving to
 * the start of the row after we reach the end) until we run out of candies.
 * The last person will receive all of our remaining candies (not necessarily
 * one more than the previous gift).
 * 
 * Return an array (of length num_people and sum candies) that represents the
 * final distribution of candies.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: candies = 7, num_people = 4
 * Output: [1,2,3,1]
 * Explanation:
 * On the first turn, ans[0] += 1, and the array is [1,0,0,0].
 * On the second turn, ans[1] += 2, and the array is [1,2,0,0].
 * On the third turn, ans[2] += 3, and the array is [1,2,3,0].
 * On the fourth turn, ans[3] += 1 (because there is only one candy left), and
 * the final array is [1,2,3,1].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: candies = 10, num_people = 3
 * Output: [5,2,3]
 * Explanation: 
 * On the first turn, ans[0] += 1, and the array is [1,0,0].
 * On the second turn, ans[1] += 2, and the array is [1,2,0].
 * On the third turn, ans[2] += 3, and the array is [1,2,3].
 * On the fourth turn, ans[0] += 4, and the final array is [5,2,3].
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= candies <= 10^9
 * 1 <= num_people <= 1000
 * 
 * 
 */

// @lc code=start
func distributeCandies(candies int, people int) []int {
	res := make([]int, people)
	
	k := root(candies)
	fmt.Println(k)

	// res[i] = (i+1) + (i+1+p) + (i+1+p*2) + ... + (i+1+n*p)
	// n = (k-i-1)/p , p = people
	candiesOf := func(i int) int {
		c0 := i+1
		n := (k - c0) / people
		cn := c0 + n*people
		return (n+1) * (c0+cn) / 2
	}

	// i<k for in case k<people
	for i := 0; i < people && i < k; i++ {
		res[i] = candiesOf(i)
	}

	// remaining
	res[k%people] += candies - k*(k+1)/2

	return res

}

func root(candies int) int {
	delta := float64(1 + 8*candies)
	return int((math.Sqrt(delta)-1)/2)
}

// @lc code=end

