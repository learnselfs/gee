// Package core @Author Bing
// @Date 2023/10/16 12:27:00
// @Desc
package core

type Group struct {
	engine *Engine
	parent *Group
	prefix string
}

func (g *Group) NewGroup(prefix string) *Group {
	group := &Group{engine: g.engine, parent: g, prefix: g.prefix + prefix}
	g.engine.groups = append(g.engine.groups, group)
	return group
}

func (g *Group) AddRoute(method, path string, handle Handle) {
	g.engine.Trie.insert(method, g.prefix+path, handle)
}

func (g *Group) GET(path string, handle Handle) {
	g.AddRoute("GET", path, handle)
}
func (g *Group) POST(path string, handle Handle) {
	g.AddRoute("POST", path, handle)
}
func (g *Group) PUT(path string, handle Handle) {
	g.AddRoute("PUT", path, handle)
}
func (g *Group) DELETE(path string, handle Handle) {
	g.AddRoute("DELETE", path, handle)
}
