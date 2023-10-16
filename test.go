// Package test @Author Bing
// @Date 2023/10/9 14:51:00
// @Desc
package main

import (
	"fmt"
	"github.com/learnselfs/gee/config"
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
	e.GET("/admin/:info", func(c *core.Context) {
		info := c.Param("info")
		c.JSON(utils.OkWithDetails(http.StatusOK, "welcome web v3 ", config.H{"info": info}))
	})
	e.GET("/home/*info", func(c *core.Context) {
		info := c.Param("info")
		c.JSON(utils.OkWithDetails(http.StatusOK, "welcome web v3 ", config.H{"info": info}))
	})
	e.POST("/home/upload", func(c *core.Context) {
		title := c.PostForm("title")
		file, err := c.PostFile("name")
		if err != nil {
			utils.Log.Println(err)
			c.JSON(utils.FailWithMsg(404, err.Error()))
		} else {
			err = c.SaveUploadFile(file, "F:\\go\\"+title)
			if err != nil {
				utils.Log.Println(err)
				c.JSON(utils.FailWithMsg(404, err.Error()))
			} else {
				c.JSON(utils.OkWithMsg(http.StatusOK, title))
			}
		}
	})
	e.POST("/home/uploads", func(c *core.Context) {
		title := c.PostForm("title")
		form, err := c.MultipartFile()
		if err != nil {
			utils.Log.Println(err)
			c.JSON(utils.FailWithMsg(404, err.Error()))
		} else {
			for _, file := range form.File["files"] {
				if err = c.SaveUploadFile(file, "f:\\go\\"+file.Filename); err != nil {
					utils.Log.Println(err)
					c.JSON(utils.FailWithMsg(404, err.Error()))
				}
				c.JSON(utils.OkWithMsg(http.StatusOK, title))
			}
		}
	})
	e.Run()
}
func testParser() {
	result := core.Parsers("/admin/info")
	fmt.Println(result)

}
func main() {
	WithV3()
	//testParser()
}
