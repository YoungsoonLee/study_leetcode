/*
 * @lc app=leetcode id=1010 lang=golang
 *
 * [1010] Pairs of Songs With Total Durations Divisible by 60
 *
 * https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/description/
 *
 * algorithms
 * Easy (46.72%)
 * Likes:    332
 * Dislikes: 32
 * Total Accepted:    23.8K
 * Total Submissions: 50.8K
 * Testcase Example:  '[30,20,150,100,40]'
 *
 * In a list of songs, the i-th song has a duration of time[i] seconds. 
 * 
 * Return the number of pairs of songs for which their total duration in
 * seconds is divisible by 60.  Formally, we want the number of indices i < j
 * with (time[i] + time[j]) % 60 == 0.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: [30,20,150,100,40]
 * Output: 3
 * Explanation: Three pairs have a total duration divisible by 60:
 * (time[0] = 30, time[2] = 150): total duration 180
 * (time[1] = 20, time[3] = 100): total duration 120
 * (time[1] = 20, time[4] = 40): total duration 60
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [60,60,60]
 * Output: 3
 * Explanation: All three pairs have a total duration of 120, which is
 * divisible by 60.
 * 
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= time.length <= 60000
 * 1 <= time[i] <= 500
 * 
 */

// @lc code=start
func numPairsDivisibleBy60(time []int) int {
	remainders := make([]int, 60)
    
    var res int
    
    for _, t := range time {
        r := t % 60
        var relaT int 
        if r > 0 {
            relaT = 60 - r
        }
        res += remainders[relaT]
        remainders[r]++
    }
    
	return res
	
	/*
	result := 0
    module := make([]int, 60)
    for _, v := range time{
        result += module[(60-v%60)%60]
        module[v%60]++
    }
	
	fmt.Println(result)
	fmt.Println(module)

    return result
	*/

	/*
	rec := [60]int{}
	
	for _, t := range time {
		rec[t%60]++
	}

	//fmt.Println(30%60)
	fmt.Println(rec)

	res := rec[0] * (rec[0]-1) / 2 // ?? why?
	res += rec[30] * (rec[30]-1) / 2 // ?? why?

	for i:=1; i<30; i++ {
		res += rec[i] * rec[60-i]
	}

	fmt.Println(res)

	return 0
	*/


	/* my solution. It's brute force
	count := 0

    for i:=0; i < len(time); i++ {
		for j:= i+1; j < len(time); j++ {
			if (time[i] + time[j]) %60 == 0 {
				count++
			}
		}
	}

	fmt.Println(count)
	return count
	*/
}
// @lc code=end

