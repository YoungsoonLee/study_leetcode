package main

type Trie struct {
	val  byte
	sons [26]*Trie
	end  int
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	size := len(word)

	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie{val: word[i]}
		}

		node = node.sons[idx]
	}
	node.end++
}

func (this *Trie) Search(word string) bool {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	if node.end > 0 {
		return true
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	return true
}

/*
type Trie struct {
	val  byte
	sons [26]*Trie
	end  int
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie{val: word[i]}
		}

		node = node.sons[idx]
	}

	node.end++
}

func (this *Trie) Search(word string) bool {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	if node.end > 0 {
		return true
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	return true
}

func main() {
	obj := Constructor()
	obj.Insert("apple")

	fmt.Println(obj.sons[0])
	fmt.Println(obj.val)
	fmt.Println(obj.end)
}
*/

/*
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
