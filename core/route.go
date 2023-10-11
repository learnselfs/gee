// Package core @Author Bing
// @Date 2023/10/9 17:41:00
// @Desc
package core

import (
	"strings"
)

type Route struct {
	Context Context           `json:"context"`
	Maps    map[string]Handle `json:"maps"`
	//Handle  Handle  `json:"handle"`
	//Prefix  string  `json:"prefix"`
}

func (r *Route) GET(path string, handle Handle) {
	r.add("GET", path, handle)
}
func (r *Route) add(method, path string, handle Handle) {
	fullPath := strings.Join([]string{method, path}, "-")
	r.Maps[fullPath] = handle
}

func NewRoute() *Route {
	return &Route{Maps: make(map[string]Handle)}
}
