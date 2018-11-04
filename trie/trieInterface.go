package trie

type TrieInterface interface {
	GetSize() int
	Find(s string) bool
	Add(s string) error
	Del(s string) error
}
