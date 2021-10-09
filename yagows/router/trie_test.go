package router

import "testing"

func TestMatchStatic(t *testing.T) {
	trie := NewTrie()

	trie.Add("/foo", "foo")
	res := trie.Match("/foo")
	if res.data[0] != "foo" {
		t.Error(res.data)
	}

	trie.Add("/bar", "bar")
	res = trie.Match("/bar")
	if res.data[0] != "bar" {
		t.Error(res.data)
	}

	trie.Add("/foo/bar", "foo/bar")
	res = trie.Match("/foo/bar")
	if res.data[0] != "foo/bar" {
		t.Error(res.data)
	}
}

func TestMatchParams(t *testing.T) {
	trie := NewTrie()

	trie.Add("/post/:id", "postId")
	res := trie.Match("/post/123")
	if res.data[0] != "postId" {
		t.Error(res.data)
	}
	if res.params["id"] != "123" {
		t.Error(res.params)
	}

	trie.Add("/post/:id/publish", "publish")
	res = trie.Match("/post/456/publish")
	if res.data[0] != "publish" {
		t.Error(res.data)
	}
	if res.params["id"] != "456" {
		t.Error(res.params)
	}

	trie.Add("/post/:id/comment/:cid", "comment")
	res = trie.Match("/post/789/comment/abc")
	if res.data[0] != "comment" {
		t.Error(res.data)
	}
	if res.params["id"] != "789" {
		t.Error(res.params)
	}
	if res.params["cid"] != "abc" {
		t.Error(res.params)
	}
}

func TestMatchNoConflict(t *testing.T) {
	trie := NewTrie()

	trie.Add("/post/:id", "postId")
	trie.Add("/:action/:name", "action")

	res := trie.Match("/post/123")
	if res.data[0] != "postId" {
		t.Error(res.data)
	}
	if res.params["id"] != "123" {
		t.Error(res.params)
	}

	res = trie.Match("/release/project")
	if res.data[0] != "action" {
		t.Error(res.data)
	}
	if res.params["action"] != "release" {
		t.Error(res.data)
	}
	if res.params["name"] != "project" {
		t.Error(res.data)
	}
}
