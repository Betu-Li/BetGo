package betgo

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by BetGo
//
// 用于定义路由映射的处理方法
type HandlerFunc func(response http.ResponseWriter, request *http.Request)

// Engine implements the interface of ServeHTTP
//
// 添加了一张路由映射表 router，key 由请求方法和路径构成，value 为用户映射的处理方法
type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path     // 解析请求的路径
	if handler, ok := engine.router[key]; ok { // 查找路由映射表，如果查到，就执行注册的处理方法
		handler(w, req)
	} else { // 如果查不到，就返回 404 NOT FOUND
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// New is the constructor of BetGo.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute registers a route with a method
//
// 注册路由映射
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
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
