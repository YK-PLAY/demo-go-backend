package auth

type RegisterReq struct {
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
}
