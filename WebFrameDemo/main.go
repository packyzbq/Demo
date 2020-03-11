package main

import (
	"WebFrameDemo/frame"
	"net/http"
)

func main() {
	e := frame.New()
	e.GET("/", func(c *frame.Context) {
		c.HTML(http.StatusOK, "<h1> hello frame <h1>")
	})
	e.GET("/hello", func(c *frame.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	e.POST("/login", func(c *frame.Context) {
		c.JSON(http.StatusOK, frame.H{
			"username": c.PostForm("username"),
		})
	})
	e.GET("/hello/:name", func(c *frame.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	e.GET("/assets/*filepath", func(c *frame.Context) {
		c.JSON(http.StatusOK, frame.H{"filepath": c.Param("filepath")})
	})

	e.Run(":9999")
}
