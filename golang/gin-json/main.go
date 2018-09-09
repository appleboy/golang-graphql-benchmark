package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Data Binding from JSON
type Data struct {
	Hello string `form:"hello" json:"hello" xml:"hello"  binding:"required"`
}

func main() {

	r := gin.New()
	r.POST("/graphql", func(c *gin.Context) {
		var json Data
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"hello": json.Hello,
			},
		})
	})

	fmt.Println("Now server is running on port 8080")
	r.Run() // listen and serve on 0.0.0.0:8080
}
