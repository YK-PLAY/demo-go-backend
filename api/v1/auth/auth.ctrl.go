package auth

import (
	"encoding/json"
	"fmt"
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

	var req RegisterReq
	json.Unmarshal(jsonData, &req)

	fmt.Printf("%+v\n", req)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
