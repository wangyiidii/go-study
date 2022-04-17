package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})
	r.POST("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})
	r.PUT("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "put",
		})
	})
	r.DELETE("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":8090")
}
