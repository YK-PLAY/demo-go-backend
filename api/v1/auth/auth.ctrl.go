package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 12)
	return string(bytes), err
}

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

	if req.Username == "" || req.Uuid == "" {
		c.JSON(200, gin.H{
			"status":   1001,
			"errorMsg": "username or uuid is empty",
		})
		return
	}

	hash, _ := hash(req.Username + req.Uuid)

	c.JSON(200, gin.H{
		"status":  0,
		"session": hash,
	})
}
