// Package utils @Author Bing
// @Date 2023/9/26 21:23:00
// @Desc
package utils

import "log"

var (
	Log = log.Default()
)

// Logger
// @Description: logger 定义
func Logger() {
	Log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.Lshortfile)
}
