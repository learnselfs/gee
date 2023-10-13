// Package core @Author Bing
// @Date 2023/10/11 13:20:00
// @Desc
package core

import (
	"github.com/learnselfs/gee/config"
	"strings"
)

func parser(s string) []string {
	temp := strings.Split(s, "/")
	var result []string
	for _, v := range temp {
		if len(v) > 0 {
			result = append(result, v)
		}
	}
	return result
}

type node struct {
	children   []*node
	handle     Handle `info:"handle,omitempty"`
	pattern    string `info:"pattern,omitempty"`
	path       string `info:"path,omitempty"`
	depth      int
	isWildcard bool `info:"isWildcard,omitempty"`
}

func (n *node) matchChild(path string) *node {
	for _, child := range n.children {
		if child.path == path || child.isWildcard {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(path string) []*node {
	var result []*node
	for _, child := range n.children {
		if child.path == path || child.isWildcard {
			result = append(result, child)
		}
	}
	return result
}

func (n *node) insert(pattern string, paths []string, handle Handle) {
	if n.depth == len(paths) || n.isWildcard {
		n.pattern = pattern
		n.handle = handle
		return
	}

	child := n.matchChild(paths[n.depth])

	if child != nil {
		child.insert(pattern, paths, handle)
	} else {
		child := &node{
			children:   make([]*node, 0),
			path:       paths[n.depth],
			depth:      n.depth + 1,
			isWildcard: paths[n.depth][0] == '*' || paths[n.depth][0] == ':',
		}
		n.children = append(n.children, child)
		child.insert(pattern, paths, handle)
	}
}

func (n *node) search(paths []string) *node {
	if n.depth == len(paths) || strings.HasPrefix(n.path, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}

	children := n.matchChildren(paths[n.depth])
	if len(children) > 0 {
		for _, child := range children {
			result := child.search(paths)
			return result
		}
	}

	return nil
}
func newNode() *node {
	return &node{children: make([]*node, 0), depth: 0}
}

type Trie struct {
	Nodes map[string]*node `json:"nodes"`
}

func (t *Trie) getNode(method string) *node {
	if _, ok := t.Nodes[method]; !ok {
		t.Nodes[method] = newNode()
	}
	return t.Nodes[method]
}

func (t *Trie) insert(method, pattern string, handle Handle) {
	n := t.getNode(method)
	paths := parser(pattern)
	if len(paths) == 0 {
		n.path = "/"
		n.pattern = "/"
		n.handle = handle
		return
	}
	n.insert(pattern, paths, handle)
}

func (t *Trie) search(method, pattern string) (*node, config.H) {
	params := make(config.H)
	n := t.getNode(method)
	userPaths := parser(pattern)
	if userPaths == nil {
		return nil, config.H{}
	}
	result := n.search(userPaths)
	if result == nil || len(result.pattern) == 0 {
		return nil, config.H{}
	}
	paths := parser(result.pattern)
	for index, path := range paths {
		if strings.HasPrefix(path, "*") {
			params[paths[index][1:]] = strings.Join(userPaths[index:], "/")
			break
		}
		if strings.HasPrefix(path, ":") {
			params[paths[index][1:]] = userPaths[index]
		}
	}
	return result, params
}

func (t *Trie) GET(pattern string, handle Handle) {
	t.insert("GET", pattern, handle)
}

func NewTrie() *Trie {
	return &Trie{Nodes: make(map[string]*node)}
}
