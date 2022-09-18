package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello goland",
	})
}

func main() {
	r := gin.Default()

	r.GET("/hello", sayHello)

	// json 返回数据
	r.GET("/json", func(c *gin.Context) {
		// 使用map
		data := map[string]interface{}{
			"name":  "Aring",
			"age":   25,
			"point": "handsome",
		}
		c.JSON(http.StatusOK, data)
	})

	// 使用struct
	r.GET("struct", func(c *gin.Context) {
		monster := struct {
			Name  string `json:"name"`
			Age   int    `json:"age"`
			Point string `json:"adfsadf"`
		}{
			"彭绮文",
			24,
			"normal",
		}
		c.JSON(http.StatusOK, monster)
	})

	err := r.Run(":9000")
	if err != nil {
		return
	}
}
