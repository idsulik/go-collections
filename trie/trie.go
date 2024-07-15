package trie

// Node represents each node in the Trie
type Node struct {
	children map[rune]*Node
	isEnd    bool
}

// Trie represents the Trie structure
type Trie struct {
	root *Node
}

// New initializes a new Trie
func New() *Trie {
	return &Trie{root: newNode()}
}

// newNode initializes a new Trie node
func newNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}

// Insert Adds a word to the Trie.
func (t *Trie) Insert(words string) {
	current := t.root
	for _, char := range words {
		if _, found := current.children[char]; !found {
			current.children[char] = newNode()
		}
		current = current.children[char]
	}
	current.isEnd = true
}

// Search searches for a word in the Trie and returns true if the word exists
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if _, found := current.children[char]; !found {
			return false
		}
		current = current.children[char]
	}
	return current.isEnd
}

// StartsWith checks if there is any word in the Trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, char := range prefix {
		if _, found := current.children[char]; !found {
			return false
		}
		current = current.children[char]
	}
	return true
}
