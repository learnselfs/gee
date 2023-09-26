// Package gee @Author Bing
// @Date 2023/9/26 21:09:00
// @Desc
package gee

import "gee/core"

func New(address, port string) *core.Engine {
	return &core.Engine{Address: address, Port: port}
}
