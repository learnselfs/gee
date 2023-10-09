// Package core @Author Bing
// @Date 2023/9/26 21:09:00
// @Desc
package core

import (
	"github.com/learnselfs/gee/utils"
	"log"
	"net/http"
)

// Engine
// @Description: 核心组件： 主要承载 官方组件 http server（实现 server 接口）
type Engine struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

func (e *Engine) handler(c *Context) {
	if c.Path == "/" {
		c.JSON(utils.Ok())
	} else {
		c.Data(404, []byte("Not Found"))
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
	log.Printf("web server listening: %s:%s", e.Address, e.Port)
	err := http.ListenAndServe(e.Address+":"+e.Port, e)
	if err != nil {
		return
	}
}
