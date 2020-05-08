/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 *
 * https://leetcode.com/problems/excel-sheet-column-number/description/
 *
 * algorithms
 * Easy (52.99%)
 * Likes:    839
 * Dislikes: 144
 * Total Accepted:    274.3K
 * Total Submissions: 510.6K
 * Testcase Example:  '"A"'
 *
 * Given a column title as appear in an Excel sheet, return its corresponding
 * column number.
 * 
 * For example:
 * 
 * 
 * ⁠   A -> 1
 * ⁠   B -> 2
 * ⁠   C -> 3
 * ⁠   ...
 * ⁠   Z -> 26
 * ⁠   AA -> 27
 * ⁠   AB -> 28 
 * ⁠   ...
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "A"
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "AB"
 * Output: 28
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "ZY"
 * Output: 701
 * 
 */

// @lc code=start
func titleToNumber(s string) int {
	res := 0
	s = "ZY"
	for i:=0; i <len(s); i++ {
		temp := int(s[i]-'A'+1)
		fmt.Println(temp)
		res = res * 26 + temp
		fmt.Println(res)
	}

	return res
    
}
// @lc code=end

