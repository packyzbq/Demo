package main

import (
	"WebFrameDemo/frame"
	"WebFrameDemo/midware"
	"fmt"
	"net/http"
)

//func main() {
//	e := frame.New()
//	e.GET("/", func(c *frame.Context) {
//		c.HTML(http.StatusOK, "<h1> hello frame <h1>")
//	})
//	e.GET("/hello", func(c *frame.Context) {
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//	})
//	e.POST("/login", func(c *frame.Context) {
//		c.JSON(http.StatusOK, frame.H{
//			"username": c.PostForm("username"),
//		})
//	})
//	e.GET("/hello/:name", func(c *frame.Context) {
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//	})
//
//	e.GET("/assets/*filepath", func(c *frame.Context) {
//		c.JSON(http.StatusOK, frame.H{"filepath": c.Param("filepath")})
//	})
//
//	e.Run(":9999")
//}

func main() {
	r := frame.New()
	r.Use(midware.Logger(), midware.Recover())
	r.GET("/index", func(c *frame.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	r.GET("/panic", func(c *frame.Context) {
		a := []int{1, 2, 3}
		fmt.Println(a[4])
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *frame.Context) {
			c.HTML(http.StatusOK, "<h1>Hello V1 Group</h1>")
		})
		v1.GET("/hello", func(c *frame.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(midware.A())
	{
		v2.GET("/hello/:name", func(c *frame.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *frame.Context) {
			c.JSON(http.StatusOK, frame.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
