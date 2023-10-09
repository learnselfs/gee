// Package utils @Author Bing
// @Date 2023/10/7 17:20:00
// @Desc
package utils

import (
	"github.com/learnselfs/gee/core"
	"net/http"
)

func msg(status string, code int, message string, data core.H) core.H {
	return core.H{"status": status, "code": code, "msg": message, "data": data}
}

func Ok() core.H {
	return msg(http.StatusText(http.StatusOK), http.StatusOK, "OK", core.H{})
}

func OkWithMsg(status string, code int, message string, data core.H) core.H {
	return msg(status, code, message, data)
}

func OkWithDetails(status string, code int, message string, data core.H) core.H {
	return msg(status, code, message, data)
}
