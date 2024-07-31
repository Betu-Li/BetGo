package betgo

import (
	"net/http"
)

// HandlerFunc defines the request handler used by BetGo
//
// 用于定义路由映射的处理方法
type HandlerFunc func(response http.ResponseWriter, request *http.Request)

// Engine implements the interface of ServeHTTP
type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// New is the constructor of BetGo.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute adds a route to the router
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
//
// 包装了 http.ListenAndServe 方法，启动 http 服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
