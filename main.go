package main

import (
	"fmt"

	"github.com/mars1385/leetcode/trie"
)

func main() {

	obj := trie.Constructor()

	obj.Insert("hotdog")
	// return True
	// obj.Search("app")     // return False
	// obj.StartsWith("app") // return True
	// obj.Insert("app")
	// obj.Search("app") // return True

	fmt.Println(obj.Search("hot"))

}
