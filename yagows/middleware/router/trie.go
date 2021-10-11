package router

import (
	"strings"
)

const WILDCARD = "*"

func isWildcard(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

type Trie struct {
	// paramName of a wildcard node, empty string for static nodes
	paramName string

	// data attached to a node, nil for intermediate nodes
	data []interface{}

	// child nodes indexed by name
	children map[string]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		children: map[string]*Trie{},
	}
}

// Add an url path to trie
func (t *Trie) Add(path string, data ...interface{}) {
	segments := strings.Split(path, "/")
	t.add(t, segments, data)
}

type MatchResult struct {
	// data attached to node
	data []interface{}

	// matched params along the path
	params map[string]string
}

// Match an url path in trie
func (t *Trie) Match(path string) *MatchResult {
	segments := strings.Split(path, "/")

	result := &MatchResult{
		params: map[string]string{},
	}
	t.match(t, segments, result)

	return result
}

func (t *Trie) add(parent *Trie, segments []string, data []interface{}) {
	if len(segments) == 0 {
		// last node, attach data
		parent.data = data
		return
	}

	// split out first segment
	seg, rest := segments[0], segments[1:]

	// for a wildcard child, add index '*' -> child, and save it's param name
	paramName := ""
	if isWildcard(seg) {
		paramName = seg[1:]
		seg = WILDCARD
	}

	child := parent.children[seg]
	// create a new child node if necessary
	if child == nil {
		child = NewTrie()
		parent.children[seg] = child
	}

	// only a single wildcard child is allowed
	// latter wildcard child overwrites previous paramName
	child.paramName = paramName

	// process remaining segments
	t.add(child, rest, data)
}

func (t *Trie) match(parent *Trie, segments []string, result *MatchResult) {
	if len(segments) == 0 {
		// finished match
		result.data = parent.data
		return
	}

	// split out first segment
	seg, rest := segments[0], segments[1:]

	child := parent.children[seg]
	if child == nil {
		// no match, try a wildcard match
		child = parent.children[WILDCARD]
		if child == nil {
			// still no match
			return
		}
		// matched wildcard, add to param mapping
		result.params[child.paramName] = seg
	}

	// process remaining segments
	t.match(child, rest, result)
}
