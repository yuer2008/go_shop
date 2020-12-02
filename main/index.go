package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-test.com/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "this is a test page")
	})
	//收货地址分组
	add := r.Group("address")
	{
		add.GET("list.html", controllers.List)
		add.GET("add.html", controllers.Add)
	}
	r.GET("/ping", func(c *gin.Context) {
		a := 10

		c.JSON(200, gin.H{
			"age":     a,
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
