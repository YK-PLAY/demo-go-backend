package auth

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//Handle error
		c.JSON(200, gin.H{
			"status": 1001,
		})
		return
	}

	print(string(jsonData))

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
