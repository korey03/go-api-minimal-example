package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./views/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
			"title":   "API Test",
			"message": "Hello world!",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":   1,
			"name": "toto",
		})
	})

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
