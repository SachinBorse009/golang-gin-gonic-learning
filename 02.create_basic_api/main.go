package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//what is api = application Programming interface

type api struct {
	Name 	string  `json:"name"`
	Email 	string  `json:"email"`
}

//create a variable for api struct type
var data api

//main function
func main() {
	r := gin.Default()  //initialize gin 

	r.GET("/get", getValues)
	r.POST("/post", postValues)
	r.PUT("/put", putValues)
	r.DELETE("/delete", deleteValues)
	
	r.Run(":4000")
}

//getvalues handler
func getValues(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

//postvalues handler
func postValues(c *gin.Context){
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message" : "Somthing is wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : data,
	})
}

//putvalues handler
func putValues(c *gin.Context){
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message":"Somthing wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

//deleteValues handler
func deleteValues(c *gin.Context){
	data = api{}
	c.JSON(http.StatusOK,gin.H{
		"message":data,
	})
}