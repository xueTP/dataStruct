package trie

import (
	"errors"
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

func (tfb TrieForByte) Find(s string) bool {
	sb := []byte(s)
	pre := tfb.data
	for _, v := range sb {
		if _, ok := pre.Next[v]; !ok {
			return false
		}
		pre = pre.Next[v]
	}
	if pre.Val == 1 {
		return true
	}
	return false
}

func (tfb *TrieForByte) Add(s string) error {
	sb := []byte(s)
	pre := tfb.data
	for _, v := range sb {
		if next, ok := pre.Next[v]; ok {
			pre = next
		} else {
			pre.Next[v] = NewTrieNode()
			pre = pre.Next[v]
		}
	}
	pre.Val = 1
	tfb.size++
	return nil
}

func (tfb *TrieForByte) Del(s string) error {
	oldSize := tfb.GetSize()
	tfb.del(tfb.data, s)
	if oldSize == tfb.GetSize() {
		return errors.New("this trie not have this word")
	}
	return nil
}

func (tfb *TrieForByte) del(node *TrieNode, s string) bool {
	if len(s) == 0 {
		tfb.size--
		node.Val = 0
		return true
	}

	if next, ok := node.Next[byte(s[0])]; ok {
		bo := tfb.del(next, s[1:])
		if len(node.Next) == 1 && bo {
			delete(node.Next, byte(s[0]))
		} else {
			return false
		}
	}
	return true
}

func TrieForByteDemo() {
	tfb := NewTrieForByte()
	tfb.Add("abcdefg")
	fmt.Println(tfb)
	tfb.Add("abcdgrt")
	fmt.Println(tfb)
	bo := tfb.Find("abcde")
	fmt.Println(tfb, bo)
	err := tfb.Del("abcdgrt")
	fmt.Println(tfb, err)
}
