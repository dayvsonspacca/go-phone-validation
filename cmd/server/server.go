package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/ping", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to read request body"})
			return
		}
		fmt.Println("Request Body:", string(body))

		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":9000")
}
