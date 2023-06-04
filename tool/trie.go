/*******
* @Author:qingmeng
* @Description:
* @File:trie
* @Date2021/12/11
 */

package tool

var sensitiveWords = make([]interface{}, 0)

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}

//字典树节点
type TrieNode struct {
	children map[interface{}]*TrieNode
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[interface{}]*TrieNode),
		isEnd:    false,
	}
}

func (trie *Trie) Insert(word []interface{}) {
	if len(word) == 0 {
		return
	}
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

func (trie *Trie) StartsWith(prefix []interface{}) bool {
	node := trie.root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.children[prefix[i]]
		if !ok {
			return false
		}
		node = node.children[prefix[i]]
	}
	return true
}
