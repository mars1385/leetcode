package trie

// type Trie struct {
// 	items map[string]bool
// }

// func Constructor() Trie {
// 	return Trie{
// 		items: map[string]bool{},
// 	}
// }

// func (this *Trie) Insert(word string) {
// 	this.items[word] = true
// }

// func (this *Trie) Search(word string) bool {
// 	if len(this.items) <= 0 {
// 		return false
// 	}
// 	_, has := this.items[word]

// 	return has
// }

// func (this *Trie) StartsWith(prefix string) bool {

// 	if len(this.items) <= 0 {
// 		return false
// 	}

// 	for key := range this.items {

// 		if strings.HasPrefix(key, prefix) {
// 			return true
// 		}
// 	}

// 	return false

// }

type Trie struct {
	items     map[rune]*Trie
	isEndLast bool
}

func Constructor() Trie {
	return Trie{
		isEndLast: false,
		items:     make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, w := range word {

		_, has := current.items[w]

		if !has {

			current.items[w] = &Trie{
				isEndLast: false,
				items:     make(map[rune]*Trie),
			}

		}

		current = current.items[w]
	}

	current.isEndLast = true
}

func (this *Trie) Search(word string) bool {
	if len(this.items) <= 0 {
		return false
	}

	current := this
	for _, w := range word {
		_, has := current.items[w]
		if !has {
			return false
		}
		current = current.items[w]

	}

	return current.isEndLast
}

func (this *Trie) StartsWith(prefix string) bool {

	if len(this.items) <= 0 {
		return false
	}
	current := this
	for _, w := range prefix {
		_, has := current.items[w]

		if !has {
			return false
		}
		current = current.items[w]

	}

	return current != nil

}
