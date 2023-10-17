// Package core @Author Bing
// @Date 2023/10/7 17:04:00
// @Desc
package core

import (
	"encoding/json"
	"fmt"
	"github.com/learnselfs/gee/config"
	"github.com/learnselfs/gee/utils"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type (
	// Context
	// @Description: request response Context management
	Context struct {
		//request
		r                  *http.Request          `info:"request"`
		path               string                 `info:"path"`
		method             string                 `info:"method"`
		params             map[string]interface{} `info:"params"`
		maxMultipartMemory int64                  `info:"maxMultipartMemory"`
		// response
		w         http.ResponseWriter `info:"response"`
		stateCode int                 `info:"stateCode"`
	}
)

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		r:                  r,
		method:             r.Method,
		path:               r.URL.Path,
		w:                  w,
		params:             make(config.H, 0),
		maxMultipartMemory: 8 << 20,
	}
}

func (c *Context) SetHeader(rType string) {
	c.w.Header().Set("Content-Type", rType)
}

func (c *Context) JSON(obj config.H) {
	utils.Log.Println(c.method, c.r.URL.Path, obj)
	c.SetHeader("application/json")
	buf := json.NewEncoder(c.w)
	err := buf.Encode(obj)
	if err != nil {
		utils.Log.Fatalln(err)
	}
}

func (c *Context) HTML(code int, html string) {
	c.stateCode = code
	c.SetHeader("text/html")
	_, i := c.w.Write([]byte(html))
	if i != nil {
		return
	}
}

func (c *Context) String(code int, format string, data ...config.H) {
	c.SetHeader("text/plain")
	c.stateCode = code
	_, i := c.w.Write([]byte(fmt.Sprintf(format, data)))
	if i != nil {
		return
	}
}

func (c *Context) Data(code int, data []byte) {
	c.stateCode = code
	_, err := c.w.Write(data)
	if err != nil {
		return
	}
}

func (c *Context) Param(key string) any {
	return c.params[key]
}

func (c *Context) PostForm(key string) string {
	return c.r.PostFormValue(key)
}

func (c *Context) PostFile(name string) (*multipart.FileHeader, error) {
	if c.r.MultipartForm == nil {
		if err := c.r.ParseMultipartForm(c.maxMultipartMemory); err != nil {
			return nil, err
		}
	}
	_, fileHeader, err := c.r.FormFile(name)
	if err != nil {
		return nil, err
	}
	return fileHeader, nil
}

func (c *Context) MultipartFile() (*multipart.Form, error) {
	err := c.r.ParseMultipartForm(c.maxMultipartMemory)
	return c.r.MultipartForm, err
}

func (c *Context) SaveUploadFile(fileHeader *multipart.FileHeader, fp string) error {

	source, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer source.Close()

	err = os.MkdirAll(filepath.Dir(fp), 0750)
	if err != nil {
		return err
	}
	destination, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err

}
