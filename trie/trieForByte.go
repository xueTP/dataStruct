package trie

import (
	"fmt"
)

type TrieForByte struct {
	data *TrieNode
	size int
}

func (tfb TrieForByte) String() string {
	return tfb.string(tfb.data)
}

func (tfb TrieForByte) string(node *TrieNode) string {
	var s string = "("
	if len(node.Next) != 0 {
		for i, v := range node.Next {
			s += string(i)
			s += tfb.string(v)
		}
	}
	s += ")"
	return s
}

func NewTrieForByte() *TrieForByte {
	return &TrieForByte{
		data: NewTrieNode(),
	}
}

func (tfb TrieForByte) GetSize() int {
	return tfb.size
}

func (*TrieForByte) Find(s string) bool {
	panic("implement me")
}

func (tfb *TrieForByte) Add(s string) error {
	sb := []byte(s)
	pre := tfb.data
	for _, v := range sb {
		if next, ok := pre.Next[v]; ok {
			pre = next
		}else {
			pre.Next[v] = NewTrieNode()
			pre = pre.Next[v]
		}
	}
	pre.Val = 1
	tfb.size ++
	return nil
}

func (*TrieForByte) Del(s string) error {
	panic("implement me")
}

func TrieForByteDemo() {
	tfb := NewTrieForByte()
	tfb.Add("abcdefg")
	fmt.Println(tfb)
	tfb.Add("abcdgrt")
	fmt.Println(tfb)
}
