package router

import "testing"

func TestMatchStatic(t *testing.T) {
	trie := NewTrie()

	trie.Add("/foo", "foo")
	trie.Add("/bar", "bar")
	trie.Add("/foo/bar", "foo/bar")

	m := trie.Match("/foo")
	if m.data[0] != "foo" {
		t.Error(m.data)
	}

	m = trie.Match("/bar")
	if m.data[0] != "bar" {
		t.Error(m.data)
	}

	m = trie.Match("/foo/bar")
	if m.data[0] != "foo/bar" {
		t.Error(m.data)
	}

	m = trie.Match("/bar/foo")
	if m.data != nil {
		t.Error(m.data)
	}
}

func TestMatchParams(t *testing.T) {
	trie := NewTrie()

	trie.Add("/post/:id", "postId")
	trie.Add("/post/:id/publish", "publish")
	trie.Add("/post/:id/comment/:cid", "comment")

	m := trie.Match("/post/123")
	if m.data[0] != "postId" {
		t.Error(m.data)
	}
	if m.params["id"] != "123" {
		t.Error(m.params)
	}

	m = trie.Match("/post/456/publish")
	if m.data[0] != "publish" {
		t.Error(m.data)
	}
	if m.params["id"] != "456" {
		t.Error(m.params)
	}

	m = trie.Match("/post/789/comment/abc")
	if m.data[0] != "comment" {
		t.Error(m.data)
	}
	if m.params["id"] != "789" {
		t.Error(m.params)
	}
	if m.params["cid"] != "abc" {
		t.Error(m.params)
	}
}

func TestMatchNoConflict(t *testing.T) {
	trie := NewTrie()

	trie.Add("/post/:id", "/post/:id")
	trie.Add("/:action/id", "/:action/id")

	m := trie.Match("/post/id")
	if m.data[0] != "/post/:id" {
		t.Error(m.data)
	}
	if m.params["id"] != "id" {
		t.Error(m.params)
	}

	m = trie.Match("/action/id")
	if m.data[0] != "/:action/id" {
		t.Error(m.data)
	}
	if m.params["action"] != "action" {
		t.Error(m.params)
	}
}

func TestMatchConflict(t *testing.T) {
	trie := NewTrie()

	trie.Add("/book/:id/title", "/book/:id/title")
	trie.Add("/book/:author/age", "/book/:author/age")

	m := trie.Match("/book/id/title")
	if m.data[0] != "/book/:id/title" {
		t.Error(m.data)
	}
	// later wildcard param name takes effect
	if m.params["author"] != "id" {
		t.Error(m.params)
	}

	m = trie.Match("/book/author/age")
	if m.data[0] != "/book/:author/age" {
		t.Error(m.data)
	}
	if m.params["author"] != "author" {
		t.Error(m.params)
	}
}
