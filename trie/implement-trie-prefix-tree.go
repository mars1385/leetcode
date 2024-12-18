package trie

import (
	"strings"
)

type Trie struct {
	items map[string]bool
}

func Constructor() Trie {
	return Trie{
		items: map[string]bool{},
	}
}

func (this *Trie) Insert(word string) {
	this.items[word] = true
}

func (this *Trie) Search(word string) bool {
	if len(this.items) <= 0 {
		return false
	}
	_, has := this.items[word]

	return has
}

func (this *Trie) StartsWith(prefix string) bool {

	if len(this.items) <= 0 {
		return false
	}

	for key := range this.items {

		if strings.HasPrefix(key, prefix) {
			return true
		}
	}

	return false

}
