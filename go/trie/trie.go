package main

// This package implements very simple trie-like data structure for storing,
// retrieval and querying of strings using longest common prefix chaining.

import (
	"fmt"
	"strings"
)

// Given two strings s1 and s2 find the length and value of their common prefix
// if any. In case there's no common prefix it returns 0, "".
func commonPrefix(s1, s2 string) (int, string) {
	l := len(s1)
	if len(s2) < l {
		l = len(s2)
	}

	n := 0
	for i := 0; i < l; i++ {
		if s1[i] == s2[i] {
			n++
		} else {
			break
		}
	}
	return n, s1[:n]
}

type Trie struct {
	leaf bool
	ch map[string]*Trie
}

func New() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(s string) {
	if t.ch == nil {
		t.ch = make(map[string]*Trie)
	}

	prefixLen := 0
	prefix := ""
	origKey := "" // original key that shares longest common prefix 
	for k, _ := range(t.ch) {
		n, p := commonPrefix(s, k)
		// n > 0 (instead of prefixLen) and break is enough as by design we'll
		// never have multiple conflicting prefixes.
		if n > 0 {
			prefixLen = n
			prefix = p
			origKey = k
			break
		}
	}

	// If there's no key sharing common prefix just add a new key.
	if prefixLen == 0 {
		t.ch[s] = nil

	} else {
		// When turning existing string terminal to a node, set leaf flag to
		// true.
		v, ok := t.ch[prefix]
		if v == nil {
			t.ch[prefix] = &Trie{
				leaf: ok,
			}
		}
		// Consider cases where word terminating in middle of existing trie is
		// added.
		if prefix == s {
			t.ch[prefix].leaf = true
		}

		// If prefix is only part of origKey we need to cut the origKey in
		// prefix:suffix and move suffix under the newly made prefix key.
		if prefix != origKey {
			t.ch[prefix].ch = make(map[string]*Trie)
			newKey := origKey[prefixLen:]
			t.ch[prefix].ch[newKey] = t.ch[origKey]
			delete(t.ch, origKey)
		}

		if s[prefixLen:] != "" {
			t.ch[prefix].Insert(s[prefixLen:])
		}
	}

}

func (t *Trie) Has(s string) bool {
	if t.ch == nil { return false }

	// To avoid range loop below try to lookup s directly in t.ch.
	v, ok := t.ch[s]
	if ok == true && (v == nil || v.leaf) {
		fmt.Println("direct hit")
		return true
	}

	sn := len(s)
	for k, v := range(t.ch) {
		n, _ := commonPrefix(s, k)
		if n == 0 {
			continue
		}
		// If length of common prefix is length of s then we have match.
		if sn == n {
			// Match can still be only intermediate node, check if it is leaf.
			return (v == nil || v.leaf == true)
		} else {
			return v.Has(s[n:])
		}
	}
	return false
}

func (t *Trie) All(prefix string) {
	for k, v := range(t.ch) {
		if v == nil || v.leaf {
			fmt.Println(prefix + k)
		}
		if v != nil {
			v.All(prefix + k)
		}
	}
}

func (t *Trie) Print(s int) {
	for k, v := range(t.ch) {
		x := false
		if v != nil { x = v.leaf } 
		fmt.Println(strings.Repeat(" ", s) + "[" + k + "]", x)
		if v != nil {
			v.Print(s+1)
		}
	}
}

func main() {
	t := New()
	//t *Trie

	t.Insert("abcda")
	t.Insert("abcdd")
	t.Insert("abcdde")
	t.Insert("abcddef")
	t.Insert("abcd")
	t.Insert("abd")
	t.Insert("foo")
	t.Insert("bar")
	t.Insert("foobar")
	t.Insert("fo")
	t.Print(0)
	fmt.Println(t.Has("foo"))
	fmt.Println(t.Has("abcd"))
	t.All("")
}
