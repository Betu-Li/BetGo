package betgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H betgo.H 作为别名，在构建 JSON 数据时使用，显得更简洁
type H map[string]interface{}

type context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string

	// response info
	StatusCode int
}

// newContext creates a new context object.
func newContext(w http.ResponseWriter, req *http.Request) *context {
	return &context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm gets a form value.
func (c *context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Status sets the HTTP response code.
func (c *context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader sets the HTTP header.
func (c *context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// String sets the string context.
func (c *context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON sets the JSON context.
func (c *context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	// json.NewEncoder(c.Writer).Encode(obj)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data sets the data context.
func (c *context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML sets the HTML context.
func (c *context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
