package auth

import commonapi "github.com/YK-PLAN/demo-go-backend/common/api"

type RegisterReq struct {
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
}

type RegisterAuthReq struct {
	Username   string `json:"username"`
	AuthNumber string `json:"authNumber"`
}

type RegisterAuthRes struct {
	commonapi.Response
	Token string `json:"token"`
}

type AuthUser struct {
	Username     string
	Uuid         string
	RandomNumber string
	UserId       string
}
