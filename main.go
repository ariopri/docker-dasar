package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", HealtCheckHandler)
	r.Run(":8080")
}

// HealtCheckHandler using gin
func HealtCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
