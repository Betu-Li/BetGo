package main

import (
	"betgo"
	"net/http"
)

func main() {
	r := betgo.New()
	r.GET("/", func(c *betgo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *betgo.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *betgo.Context) {
		c.JSON(http.StatusOK, betgo.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
