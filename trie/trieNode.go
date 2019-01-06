package trie

// 用于字典树的节点
type TrieNode struct {
	Val  byte
	Next map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Next: make(map[byte]*TrieNode)}
}

func NewTrieNodeWithVal(val byte) *TrieNode {
	return &TrieNode{Val: val, Next: make(map[byte]*TrieNode)}
}
