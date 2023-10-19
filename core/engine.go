// Package core @Author Bing
// @Date 2023/9/26 21:09:00
// @Desc
package core

import (
	"github.com/learnselfs/gee/utils"
	"log"
	"net/http"
	"strings"
)

// Engine
// @Description: 核心组件： 主要承载 官方组件 http server（实现 server 接口）
type Engine struct {
	//*Route  `json:"route"`
	*Trie `json:"trie"`
	*Group
	Address string   `json:"address"`
	Port    string   `json:"port"`
	groups  []*Group `info:"groups"`
}

// handler
// @Description: server's processor
// @receiver e
// @param c
func (e *Engine) handler(c *Context) {
	n, params := e.search(c.method, c.path)
	if n != nil && n.handle != nil {
		c.params = params
		c.middleware = append(c.middleware, n.handle)
		c.Next()
	} else {
		c.JSON(utils.Fail())
	}

}

// ServeHTTP
// @Description: 实现 http Handler 接口
// @param w http.ResponseWriter
// @param r *http.Request
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	for _, group := range e.groups {
		if strings.HasPrefix(c.path, group.prefix) {
			c.middleware = append(c.middleware, group.middleware...)
		}
	}
	e.handler(c)
}

// Run
// @Description: 启动函数
func (e *Engine) Run() {
	log.Printf("V3 web server listening: %s:%s", e.Address, e.Port)
	err := http.ListenAndServe(e.Address+":"+e.Port, e)
	if err != nil {
		return
	}
}

func NewEngine(address, port string) *Engine {
	engine := &Engine{Address: address, Port: port, groups: make([]*Group, 0), Trie: NewTrie()}
	engine.Group = &Group{engine: engine}
	engine.groups = append(engine.groups, engine.Group)
	return engine
}
