package auth

import (
	"fmt"
	"math/rand"
	"time"

	commonapi "github.com/YK-PLAN/demo-go-backend/common/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var digitalRunes = []rune("0123456789")
var userMap = make(map[string]AuthUser)

func randomNumber(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = digitalRunes[rand.Intn(len(digitalRunes))]
	}

	return string(b)
}

func createJwt() (string, string, error) {
	// make uuid
	userId, err := uuid.NewV4()
	if err != nil {
		return "", "", err
	}

	userIdString := userId.String()

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"userId": userIdString,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
		"iat":    time.Now().Unix(),
	}

	keyBytes := []byte("hello")
	tokenString, err2 := token.SignedString(keyBytes)
	if err2 != nil {
		fmt.Printf("error: %s\n", err2.Error())
		return "", "", err2
	}

	return tokenString, userIdString, nil
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
			jwt, userId, err := createJwt()
			if err != nil {
				commonapi.MakeResponse(c, commonapi.INTERNAL_ERROR, "Invalid auth number")
			} else {
				val.UserId = userId
				res := RegisterAuthRes{Token: jwt}
				commonapi.MakeResponseWithBody(c, commonapi.OK, "", res)
			}
		} else {
			commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "Invalid auth number")
		}
	} else {
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "Unregister username")
	}
}

func parseToken(c *gin.Context) {
	reqMap := make(map[string]string)
	err := commonapi.ReadRequest(c, &reqMap)
	if err != nil {
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "read json error")
		return
	}

	fmt.Printf("[parseToken]Req: %+v\n", reqMap)

	tokenString := reqMap["token"]
	fmt.Printf("TokenString: %s\n", tokenString)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("hello"), nil
	})

	fmt.Printf("Token: %+v\n", token)

	claims := token.Claims.(jwt.MapClaims)
	fmt.Printf("Hello %s\n", claims["userId"])

	commonapi.MakeResponse(c, commonapi.OK, "")
}
