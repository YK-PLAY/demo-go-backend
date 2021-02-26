package auth

import commonapi "github.com/YK-PLAN/demo-go-backend/common/api"

type RegisterReq struct {
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
}

type RegisterRes struct {
	commonapi.Response
	SessionKey string `json:"sessionKey"`
}
