/*
 * @lc app=leetcode id=599 lang=golang
 *
 * [599] Minimum Index Sum of Two Lists
 *
 * https://leetcode.com/problems/minimum-index-sum-of-two-lists/description/
 *
 * algorithms
 * Easy (48.63%)
 * Likes:    403
 * Dislikes: 156
 * Total Accepted:    64.7K
 * Total Submissions: 133K
 * Testcase Example:  '["Shogun","Tapioca Express","Burger King","KFC"]\n["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]'
 *
 * 
 * Suppose Andy and Doris want to choose a restaurant for dinner, and they both
 * have a list of favorite restaurants represented by strings. 
 * 
 * 
 * You need to help them find out their common interest with the least list
 * index sum. If there is a choice tie between answers, output all of them with
 * no order requirement. You could assume there always exists an answer.
 * 
 * 
 * 
 * Example 1:
 * 
 * Input:
 * ["Shogun", "Tapioca Express", "Burger King", "KFC"]
 * ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse",
 * "Shogun"]
 * Output: ["Shogun"]
 * Explanation: The only restaurant they both like is "Shogun".
 * 
 * 
 * 
 * Example 2:
 * 
 * Input:
 * ["Shogun", "Tapioca Express", "Burger King", "KFC"]
 * ["KFC", "Shogun", "Burger King"]
 * Output: ["Shogun"]
 * Explanation: The restaurant they both like and have the least index sum is
 * "Shogun" with index sum 1 (0+1).
 * 
 * 
 * 
 * 
 * Note:
 * 
 * The length of both lists will be in the range of [1, 1000].
 * The length of strings in both lists will be in the range of [1, 30].
 * The index is starting from 0 to the list length minus 1.
 * No duplicates in both lists.
 * 
 * 
 */

/* 
func findRestaurant(list1 []string, list2 []string) []string {
	
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}

	l1m := make(map[string]int, len(list1))
	l2m := make(map[string]int, len(list2))

	for i, w := range list1 {
		l1m[w] = i
	}

	for i, w := range list2 {
		l2m[w] = i
	}

	fmt.Println(l1m, l2m)
	sum := len(list1) + len(list2)

	for k, v := range l1m {
		v2, ok := l2m[k]
		if ok {
			sum = min(sum, v+v2)
		}
	}

	fmt.Println(sum, list2[sum])
	return []string{list2[sum]}

}
*/

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}

	m2 := make(map[string]int, len(list2))

	for i,v := range list2 {
		m2[v] = i
	}

	//fmt.Println(m2)

	min := len(list1) + len(list2)
	res := make([]string, 0, 1000)

	/*
	for i, v1 := range list1 {
		if j, ok := m2[v1]; ok {
			if i+j < min {
				min = i+j
				res = append(res, v1)
			}
		}
	}


	fmt.Println(res[len(res)-1])
	return res
	*/
	for i, r := range list1 {
		if j, ok := m2[r]; ok {
			//fmt.Println(min, i, j, i+j, res)
			if min == i+j {
				res = append(res, r)
			}
			if min > i+j {
				min = i + j
				res = append(res[len(res):], r) // !!!!
			}
		}
	}

	fmt.Println(res)
	return res

}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

