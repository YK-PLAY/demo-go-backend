package auth

import (
	"fmt"
	"math/rand"

	commonapi "github.com/YK-PLAN/demo-go-backend/common/api"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var digitalRunes = []rune("0123456789")
var userMap = make(map[string]AuthUser)

func hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 12)
	return string(bytes), err
}

func randomNumber(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = digitalRunes[rand.Intn(len(digitalRunes))]
	}

	return string(b)
}

func register(c *gin.Context) {
	var req RegisterReq
	err := commonapi.ReadRequest(c, &req)
	if err != nil {
		//Handle error
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "read json error")
		return
	}

	fmt.Printf("[register]Req: %+v\n", req)

	if req.Username == "" || req.Uuid == "" {
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "username or uuid is empty")
		return
	}

	randomNumber := randomNumber(6)
	fmt.Printf("Auth number: %s\n", randomNumber)
	userMap[req.Username] = AuthUser{Username: req.Username, Uuid: req.Uuid, RandomNumber: randomNumber}

	commonapi.MakeResponse(c, commonapi.OK, "")
}

func registerAuth(c *gin.Context) {
	var req RegisterAuthReq
	err := commonapi.ReadRequest(c, &req)
	if err != nil {
		//Handle error
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "read json error")
		return
	}

	fmt.Printf("[registerAuth]Req: %+v\n", req)

	if req.Username == "" || req.AuthNumber == "" {
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "username or auth number is empty")
		return
	}

	if val, ok := userMap[req.Username]; ok {
		if val.RandomNumber == req.AuthNumber {
			hash, err := hash(val.Username + val.Uuid)
			if err != nil {
				commonapi.MakeResponse(c, commonapi.INTERNAL_ERROR, "Invalid auth number")
			} else {
				res := RegisterAuthRes{SessionKey: hash}
				commonapi.MakeResponseWithBody(c, commonapi.OK, "", res)
			}
		} else {
			commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "Invalid auth number")
		}
	} else {
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "Unregister username")
	}
}
