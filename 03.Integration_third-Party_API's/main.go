package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// main function
func main() {
	r := gin.Default() //initialize gin

	r.GET("/get", getValues)
	r.Run(":4040")
}

// API url
var url = "http://date.jsontest.com/"

func getValues(c *gin.Context) {
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	//Close response body
	defer resp.Body.Close() //use for date bridge issue

	//red all response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	var target map[string]interface{}
	err = json.Unmarshal(data, &target)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		c.JSON(http.StatusOK, gin.H{
			"message": target,
		})
	}
}
