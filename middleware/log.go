// Package middleware @Author Bing
// @Date 2023/10/17 10:56:00
// @Desc
package middleware

import (
	"fmt"
	"github.com/learnselfs/gee/core"
)

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
