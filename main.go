// Package gee @Author Bing
// @Date 2023/9/26 21:09:00
// @Describe gee main function
package gee

import "github.com/learnselfs/gee/core"

func New(address, port string) *core.Engine {
	return core.NewEngine(address, port)
}
