// Package test @Author Bing
// @Date 2023/10/9 14:51:00
// @Desc
package main

import (
	"fmt"
	"github.com/learnselfs/gee/core"
	"github.com/learnselfs/gee/middleware"
	"github.com/learnselfs/gee/utils"
	"net/http"
	"path/filepath"
)

func WithV2() {
	e := New("localhost", "8088")
	e.Run()
}

//func WithV3() {
//	e := New("localhost", "8088")
//	e.GET("/", func(c *core.Context) {
//		c.JSON(utils.OkWithMsg(http.StatusOK, "welcome web v3 home"))
//	})
//	e.GET("/admin/:info", func(c *core.Context) {
//		info := c.Param("info")
//		c.JSON(utils.OkWithDetails(http.StatusOK, "welcome web v3 ", config.H{"info": info}))
//	})
//	e.GET("/home/*info", func(c *core.Context) {
//		info := c.Param("info")
//		c.JSON(utils.OkWithDetails(http.StatusOK, "welcome web v3 ", config.H{"info": info}))
//	})
//	e.POST("/home/upload", func(c *core.Context) {
//		title := c.PostForm("title")
//		file, err := c.PostFile("name")
//		if err != nil {
//			utils.Log.Println(err)
//			c.JSON(utils.FailWithMsg(404, err.Error()))
//		} else {
//			err = c.SaveUploadFile(file, "F:\\go\\"+title)
//			if err != nil {
//				utils.Log.Println(err)
//				c.JSON(utils.FailWithMsg(404, err.Error()))
//			} else {
//				c.JSON(utils.OkWithMsg(http.StatusOK, title))
//			}
//		}
//	})
//	e.POST("/home/uploads", func(c *core.Context) {
//		title := c.PostForm("title")
//		form, err := c.MultipartFile()
//		if err != nil {
//			utils.Log.Println(err)
//			c.JSON(utils.FailWithMsg(404, err.Error()))
//		} else {
//			for _, file := range form.File["files"] {
//				if err = c.SaveUploadFile(file, "f:\\go\\"+file.Filename); err != nil {
//					utils.Log.Println(err)
//					c.JSON(utils.FailWithMsg(404, err.Error()))
//				}
//				c.JSON(utils.OkWithMsg(http.StatusOK, title))
//			}
//		}
//	})
//	e.Run()
//}

//func WithV4() {
//	engine := New("localhost", "8088")
//	home := engine.NewGroup("/home")
//	{
//		home.GET("/index", func(c *core.Context) {
//			c.JSON(utils.OkWithMsg(http.StatusOK, "index"))
//		})
//	}
//
//	manager := engine.NewGroup("/manager")
//	{
//		manager.GET("/index", func(c *core.Context) {
//			c.JSON(utils.OkWithMsg(http.StatusOK, "manager/ index"))
//		})
//		manager.POST("/upload", func(c *core.Context) {
//			title := c.PostForm("title")
//			file, err := c.PostFile("file")
//			if err != nil {
//				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
//				return
//			}
//			err = c.SaveUploadFile(file, "f:\\go\\"+fmt.Sprintf("%s", title))
//			if err != nil {
//				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
//				return
//			}
//			c.JSON(utils.OkWithMsg(http.StatusOK, "manager/ upload"))
//		})
//
//		manager.POST("/uploads", func(c *core.Context) {
//			title := c.PostForm("title")
//			form, err := c.MultipartFile()
//			if err != nil {
//				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
//				return
//			}
//			files := form.File["files"]
//			for _, file := range files {
//				err = c.SaveUploadFile(file, "f:\\go\\"+file.Filename)
//				if err != nil {
//					c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
//					return
//				}
//			}
//			c.JSON(utils.OkWithMsg(http.StatusOK, "manager"+title))
//		})
//	}
//	engine.Run()
//}

//func WithV5() {
//
//	engine := New("localhost", "8088")
//	engine.Use(middleware.Log1())
//	home := engine.NewGroup("/home")
//	home.Use(middleware.Log4())
//	home.Use(middleware.Log5())
//	home.Use(middleware.Log6())
//	{
//		home.GET("/index", func(c *core.Context) {
//			c.JSON(utils.OkWithMsg(http.StatusOK, "index"))
//		})
//	}
//
//	manager := engine.NewGroup("/manager")
//	manager.Use(middleware.Log2())
//	manager.Use(middleware.Log3())
//	{
//		manager.GET("/index", func(c *core.Context) {
//			c.JSON(utils.OkWithMsg(http.StatusOK, "manager/ index"))
//		})
//		manager.POST("/upload", func(c *core.Context) {
//			title := c.PostForm("title")
//			file, err := c.PostFile("file")
//			if err != nil {
//				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
//				return
//			}
//			err = c.SaveUploadFile(file, "f:\\go\\"+fmt.Sprintf("%s", title))
//			if err != nil {
//				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
//				return
//			}
//			c.JSON(utils.OkWithMsg(http.StatusOK, "manager/ upload"))
//		})
//
//		manager.POST("/uploads", func(c *core.Context) {
//			title := c.PostForm("title")
//			form, err := c.MultipartFile()
//			if err != nil {
//				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
//				return
//			}
//			files := form.File["files"]
//			for _, file := range files {
//				err = c.SaveUploadFile(file, "f:\\go\\"+file.Filename)
//				if err != nil {
//					c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
//					return
//				}
//			}
//			c.JSON(utils.OkWithMsg(http.StatusOK, "manager"+title))
//		})
//	}
//	engine.Run()
//}

func WithV6() {

	engine := New("localhost", "8088")
	engine.Use(middleware.Log1())
	home := engine.NewGroup("/home")
	home.Use(middleware.Log4())
	home.Use(middleware.Log5())
	home.Use(middleware.Log6())
	{
		home.GET("/index", func(c *core.Context) {
			c.JSON(utils.OkWithMsg(http.StatusOK, "index"))
		})
	}

	manager := engine.NewGroup("/manager")
	manager.Use(middleware.Log2())
	manager.Use(middleware.Log3())
	{
		manager.GET("/index", func(c *core.Context) {
			c.JSON(utils.OkWithMsg(http.StatusOK, "manager/ index"))
		})
		manager.POST("/upload", func(c *core.Context) {
			title := c.PostForm("title")
			file, err := c.PostFile("file")
			if err != nil {
				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
				return
			}
			err = c.SaveUploadFile(file, "f:\\go\\"+fmt.Sprintf("%s", title))
			if err != nil {
				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
				return
			}
			c.JSON(utils.OkWithMsg(http.StatusOK, "manager/ upload"))
		})

		manager.POST("/uploads", func(c *core.Context) {
			title := c.PostForm("title")
			form, err := c.MultipartFile()
			if err != nil {
				c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
				return
			}
			files := form.File["files"]
			for _, file := range files {
				err = c.SaveUploadFile(file, "f:\\go\\"+file.Filename)
				if err != nil {
					c.JSON(utils.FailWithMsg(http.StatusNotFound, err.Error()))
					return
				}
			}
			c.JSON(utils.OkWithMsg(http.StatusOK, "manager"+title))
		})
	}
	// http://127.0.0.1:8088/utils/logger.go
	engine.Static("/utils")
	engine.Run()
}

func testParser() {
	result := core.Parsers("/admin/info")
	fmt.Println(result)
}

func testFileServer() {
	p, _ := filepath.Abs("./")
	path := http.Dir(p)
	FileHandle := http.FileServer(http.Dir(path))
	fileHandle := http.StripPrefix("/a/", FileHandle)
	utils.Log.Fatal(http.ListenAndServe("127.0.0.1:8088", fileHandle))
	//utils.Log.Println(path)path
}

func main() {
	WithV6()
}
