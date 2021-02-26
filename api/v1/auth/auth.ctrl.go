package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	commonapi "github.com/YK-PLAN/demo-go-backend/common/api"
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
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "read json error")
		return
	}

	var req RegisterReq
	json.Unmarshal(jsonData, &req)

	fmt.Printf("%+v\n", req)

	if req.Username == "" || req.Uuid == "" {
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "username or uuid is empty")
		return
	}

	hash, _ := hash(req.Username + req.Uuid)

	var res RegisterRes
	res.SessionKey = hash

	commonapi.MakeResponseWithBody(c, commonapi.OK, "", res)
}
