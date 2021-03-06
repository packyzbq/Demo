package frame

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router       *router
	*RouterGroup // 将engine 作为顶层 group，拥有 RouterGroup所有能力
	groups       []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.midWare = append(g.midWare, middlewares...)
}

func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.midWare...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	e.router.handle(c)
}
