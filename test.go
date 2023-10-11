// Package test @Author Bing
// @Date 2023/10/9 14:51:00
// @Desc
package main

import (
	"github.com/learnselfs/gee/core"
	"github.com/learnselfs/gee/utils"
	"net/http"
)

func WithV2() {
	e := New("localhost", "8088")
	e.Run()
}

func WithV3() {
	e := New("localhost", "8088")
	e.GET("/", func(c *core.Context) {
		c.JSON(utils.OkWithMsg(http.StatusOK, "welcome web v3 home"))
	})
	e.Run()
}

func main() {
	WithV3()
}
