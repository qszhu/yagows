package router

import (
	"strings"
)

const WILDCARD = "*"

func isWildcard(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

type node struct {
	paramName string
	isLeaf    bool
	data      []interface{}
	children  map[string]*node
}

func newNode() *node {
	return &node{
		data:     []interface{}{},
		children: map[string]*node{},
	}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

func (t *Trie) add(parent *node, segments []string, data []interface{}) {
	if len(segments) == 0 {
		parent.isLeaf = true
		parent.data = data
		return
	}

	seg, rest := segments[0], segments[1:]

	paramName := ""
	if isWildcard(seg) {
		paramName = seg[1:]
		seg = WILDCARD
	}

	child := parent.children[seg]
	if child == nil {
		child = newNode()
		if paramName != "" {
			child.paramName = paramName
		}
		parent.children[seg] = child
	}

	t.add(child, rest, data)
}

func (t *Trie) Add(url string, data ...interface{}) {
	segments := strings.Split(url, "/")
	t.add(t.root, segments, data)
}

type MatchResult struct {
	data   []interface{}
	params map[string]string
}

func (t *Trie) match(parent *node, segments []string, result *MatchResult) {
	if len(segments) == 0 {
		if parent.isLeaf {
			result.data = parent.data
		}
		return
	}

	seg, rest := segments[0], segments[1:]

	child := parent.children[seg]
	if child == nil {
		child = parent.children[WILDCARD]
		if child == nil {
			return
		}
		result.params[child.paramName] = seg
	}

	t.match(child, rest, result)
}

func (t *Trie) Match(url string) *MatchResult {
	segments := strings.Split(url, "/")

	result := &MatchResult{
		params: map[string]string{},
	}
	t.match(t.root, segments, result)

	return result
}
