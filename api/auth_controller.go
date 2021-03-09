package api

import (
	"log"
	"math/rand"

	commonapi "github.com/YK-PLAN/demo-go-backend/common/api"
	"github.com/YK-PLAN/demo-go-backend/common/db/model"
	"github.com/YK-PLAN/demo-go-backend/common/db/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthControllerV1 struct {
	db *gorm.DB
}

type RegisterReq struct {
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
}

type RegisterRes struct {
	commonapi.Response
	Action  string `json:"action"`
	Message string `json:"uuid"`
}

var digitalRunes = []rune("0123456789")

func NewAuthControllerV1(db *gorm.DB) AuthControllerV1 {
	auth := AuthControllerV1{db: db}
	return auth
}

func randomNumber(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = digitalRunes[rand.Intn(len(digitalRunes))]
	}

	return string(b)
}

func (auth *AuthControllerV1) register(c *gin.Context) {
	var req RegisterReq
	err := commonapi.ReadRequest(c, &req)
	if err != nil {
		//Handle error
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "read json error")
		return
	}

	log.Printf("[register]Req: %+v\n", req)

	if req.Username == "" || req.Uuid == "" {
		commonapi.MakeResponse(c, commonapi.INVALID_PARAM, "username or uuid is empty")
		return
	}

	randomNumber := randomNumber(6)
	log.Printf("[%s]Auth number: %s\n", req.Username, randomNumber)

	user := repository.GetUserByCellphone(auth.db, req.Username)
	if user == (model.User{}) {
		res := RegisterRes{
			Action:  "SMS_AUTH",
			Message: "New user, Please authenticate your id",
		}
		commonapi.MakeResponseWithBody(c, commonapi.OK, "", res)
		return
	}

	user = repository.GetUserByCellphoneAndUuid(auth.db, req.Username, req.Uuid)
	if user == (model.User{}) {
		res := RegisterRes{
			Action:  "SMS_AUTH",
			Message: "Unauthorized uuid, Please re-authenticate your id",
		}
		commonapi.MakeResponseWithBody(c, commonapi.OK, "", res)
		return
	}

	res := RegisterRes{
		Action:  "LOGIN",
		Message: "Registered user. Please login",
	}
	commonapi.MakeResponseWithBody(c, commonapi.OK, "", res)
}
