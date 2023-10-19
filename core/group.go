// Package core @Author Bing
// @Date 2023/10/16 12:27:00
// @Desc
package core

import (
	"fmt"
	"github.com/learnselfs/gee/utils"
	"net/http"
	"path/filepath"
)

type Group struct {
	engine     *Engine
	parent     *Group
	prefix     string
	middleware []Handle
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

func (g *Group) Use(handle Handle) {
	g.middleware = append(g.middleware, handle)
}

func (g *Group) createHandler(path string) Handle {
	absPath, _ := filepath.Abs("." + path)
	httpPath := http.Dir(absPath)
	filePath := http.StripPrefix(path, http.FileServer(httpPath))
	return func(c *Context) {
		file := c.Param("file")
		if _, err := httpPath.Open(fmt.Sprintf("%s", file)); err != nil {
			c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
		}
		filePath.ServeHTTP(c.w, c.r)
	}
}

func (g *Group) Static(path string) {
	handle := g.createHandler(path)
	g.GET(path+"/*file", handle)
}
