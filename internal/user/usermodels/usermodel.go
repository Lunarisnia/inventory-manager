package usermodels

type UserLoginCredential struct {
	NIS      string `json:"nis"`
	Password string `json:"password"`
}
