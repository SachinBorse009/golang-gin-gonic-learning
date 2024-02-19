package main

//import gin package and http package
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create main function
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from gin...!!",
		})
	})

	r.Run()
}
