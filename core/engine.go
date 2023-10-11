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
	*Route  `json:"route"`
	Address string `json:"address"`
	Port    string `json:"port"`
}

// handler
// @Description: server's processor
// @receiver e
// @param c
func (e *Engine) handler(c *Context) {
	fullPath := strings.Join([]string{c.Method, c.Path}, "-")
	if handle, ok := e.Maps[fullPath]; ok {
		handle(c)
	} else {
		c.JSON(utils.Fail())
	}
}

// ServeHTTP
// @Description: 实现 http Handler 接口
// @param w http.ResponseWriter
// @param r *http.Request
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
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
