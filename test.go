// Package test @Author Bing
// @Date 2023/10/9 14:51:00
// @Desc
package main

func WithV2() {
	e := New("localhost", "8088")
	e.Run()
}

func main() {
	WithV2()
}
