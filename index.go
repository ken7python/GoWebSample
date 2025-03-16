package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob(("templates/*"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello, %s", name)
	})
	r.POST("/user", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(200, "Hello, "+name)
	})
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin",
			"status":  200,
		})
	})
	r.Run(":8080")
}
