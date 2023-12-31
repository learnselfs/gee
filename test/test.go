// Package test @Author Bing
// @Date 2023/10/9 14:51:00
// @Desc
package test

import (
	"fmt"
	"github.com/learnselfs/gee"
	"github.com/learnselfs/gee/config"
	"github.com/learnselfs/gee/core"
	"github.com/learnselfs/gee/middleware"
	"github.com/learnselfs/gee/utils"
	"net/http"
	"path/filepath"
	"time"
)

func WithV2() {
	e := gee.New("localhost", "8088")
	e.Run()
}

func WithV3() {
	e := gee.New("localhost", "8088")
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

func WithV4() {
	engine := gee.New("localhost", "8088")
	home := engine.NewGroup("/home")
	{
		home.GET("/index", func(c *core.Context) {
			c.JSON(utils.OkWithMsg(http.StatusOK, "index"))
		})
	}

	manager := engine.NewGroup("/manager")
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
	engine.Run()
}

func WithV5() {

	engine := gee.New("localhost", "8088")
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
	engine.Run()
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

/* v6 middleware */

func Log1() core.Handle {
	return func(c *core.Context) {
		fmt.Println("1111")
		c.Next()
		fmt.Println("1111")
	}
}
func Log2() core.Handle {
	return func(c *core.Context) {
		fmt.Println("2222")
		c.Next()
		fmt.Println("2222")
	}
}
func Log3() core.Handle {
	return func(c *core.Context) {
		fmt.Println(3333)
		c.Next()
		fmt.Println(3333)
	}
}

func Log4() core.Handle {
	return func(c *core.Context) {
		fmt.Println(4444)
		c.Next()
		fmt.Println(4444)
	}
}

func Log5() core.Handle {
	return func(c *core.Context) {
		fmt.Println(5555)
		c.Next()
		fmt.Println(5555)
	}
}

func Log6() core.Handle {
	return func(c *core.Context) {
		fmt.Println(6666)
		c.Next()
		fmt.Println(6666)
	}
}

func WithV6() {
	type user struct {
		Id       int
		Username string
		Password string
	}
	engine := gee.New("localhost", "8088")

	// http://127.0.0.1:8088/utils/logger.go
	engine.Static("/static")
	engine.LoadHtml("templates/*")

	engine.Use(Log1())
	home := engine.NewGroup("/home")
	home.Use(Log4())
	home.Use(Log5())
	home.Use(Log6())
	{
		home.GET("/index", func(c *core.Context) {
			user1 := &user{Id: 1, Username: "user1", Password: "pass1"}
			user2 := &user{Id: 2, Username: "user2", Password: "pass2"}
			c.HTML(http.StatusOK,
				"index.html",
				config.H{"title": "index information", "users": [2]*user{user1, user2}})
		})

		home.GET("/base", func(c *core.Context) {
			c.HTML(http.StatusOK,
				"base.hktml",
				config.H{"title": "index information", "user": "admin"})

		})
	}

	manager := engine.NewGroup("/manager")
	manager.Use(Log2())
	manager.Use(Log3())
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
