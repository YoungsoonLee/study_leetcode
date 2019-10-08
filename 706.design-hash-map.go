/*
 * @lc app=leetcode id=706 lang=golang
 *
 * [706] Design HashMap
 *
 * https://leetcode.com/problems/design-hashmap/description/
 *
 * algorithms
 * Easy (57.11%)
 * Likes:    433
 * Dislikes: 71
 * Total Accepted:    49K
 * Total Submissions: 85.4K
 * Testcase Example:  '["MyHashMap","put","put","get","get","put","get", "remove", "get"]\n[[],[1,1],[2,2],[1],[3],[2,1],[2],[2],[2]]'
 *
 * Design a HashMap without using any built-in hash table libraries.
 *
 * To be specific, your design should include these functions:
 *
 *
 * put(key, value) : Insert a (key, value) pair into the HashMap. If the value
 * already exists in the HashMap, update the value.
 * get(key): Returns the value to which the specified key is mapped, or -1 if
 * this map contains no mapping for the key.
 * remove(key) : Remove the mapping for the value key if this map contains the
 * mapping for the key.
 *
 *
 *
 * Example:
 *
 *
 * MyHashMap hashMap = new MyHashMap();
 * hashMap.put(1, 1);
 * hashMap.put(2, 2);
 * hashMap.get(1);            // returns 1
 * hashMap.get(3);            // returns -1 (not found)
 * hashMap.put(2, 1);          // update the existing value
 * hashMap.get(2);            // returns 1
 * hashMap.remove(2);          // remove the mapping for 2
 * hashMap.get(2);            // returns -1 (not found)
 *
 *
 *
 * Note:
 *
 *
 * All keys and values will be in the range of [0, 1000000].
 * The number of operations will be in the range of [1, 10000].
 * Please do not use the built-in HashMap library.
 *
 *
 */

// @lc code=start
type MyHashMap struct {
	table []int
}

func Hash(key int) int {
	A := 1000001
	B := 3571

	h := (key + A) % B

	return h
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{table: make([]int, 1000001)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	h := Hash(key)
	this.table[h] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	h := Hash(key)

	return this.table[h]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	h := Hash(key)
	this.table[h] = 0
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

