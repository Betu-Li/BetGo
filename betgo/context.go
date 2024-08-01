package betgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H betgo.H 作为别名，在构建 JSON 数据时使用，显得更简洁
type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

// newContext creates a new Context object.
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm gets a form value.
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query gets a query parameter.
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status sets the HTTP response code.
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader sets the HTTP header.
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// String sets the string Context.
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON sets the JSON Context.
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	// json.NewEncoder(c.Writer).Encode(obj)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data sets the data Context.
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML sets the HTML Context.
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
