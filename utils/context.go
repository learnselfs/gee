// Package utils @Author Bing
// @Date 2023/10/7 17:20:00
// @Desc
package utils

import (
	"github.com/learnselfs/gee/config"
	"net/http"
)

func msg(status string, code int, message string, data config.H) config.H {
	return config.H{"status": status, "code": code, "msg": message, "data": data}
}

func Ok() config.H {
	return msg(http.StatusText(http.StatusOK), http.StatusOK, "OK", config.H{})
}

func OkWithMsg(code int, message string) config.H {
	return msg(http.StatusText(code), code, message, config.H{})
}

func OkWithDetails(code int, message string, data config.H) config.H {
	return msg(http.StatusText(code), code, message, data)
}

func Fail() config.H {
	return msg(http.StatusText(http.StatusNotFound), http.StatusNotFound, "not found page", config.H{})
}
