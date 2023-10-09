// Package core @Author Bing
// @Date 2023/10/7 17:04:00
// @Desc
package core

import (
	"encoding/json"
	"fmt"
	"github.com/learnselfs/gee/config"
	"github.com/learnselfs/gee/utils"
	"net/http"
)

type (
	// Context
	// @Description: request response context management
	Context struct {
		//request
		R      *http.Request `json:"request"`
		Path   string        `json:"path"`
		Method string        `json:"method"`
		// response
		W         http.ResponseWriter
		StateCode int
	}
)

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		R:      r,
		Method: r.Method,
		Path:   r.URL.Path,
		W:      w,
	}
}

func (c *Context) SetHeader(rType string) {
	c.W.Header().Set("Content-Type", rType)
}

func (c *Context) JSON(obj config.H) {
	utils.Log.Println(c.R.URL.Path)
	c.SetHeader("application/json")
	buf := json.NewEncoder(c.W)
	err := buf.Encode(obj)
	if err != nil {
		utils.Log.Fatalln(err)
		return
	}
}

func (c *Context) HTML(code int, html string) {
	c.StateCode = code
	c.SetHeader("text/html")
	_, i := c.W.Write([]byte(html))
	if i != nil {
		return
	}
}

func (c *Context) String(code int, format string, data ...config.H) {
	c.SetHeader("text/plain")
	c.StateCode = code
	_, i := c.W.Write([]byte(fmt.Sprintf(format, data)))
	if i != nil {
		return
	}
}

func (c *Context) Data(code int, data []byte) {
	c.StateCode = code
	_, err := c.W.Write(data)
	if err != nil {
		return
	}
}
