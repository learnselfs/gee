// Package core @Author Bing
// @Date 2023/9/26 21:09:00
// @Desc
package core

import "github.com/learnselfs/gee/utils"

// init
// @Description: 插件注册
func init() {
	utils.Logger()
}

func Parsers(s string) []string {
	return parser(s)
}
