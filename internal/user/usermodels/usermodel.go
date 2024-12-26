package usermodels

type UserLoginCredential struct {
	NIS      string `json:"nis"`
	Password string `json:"password"`
}

type GetUserInfoResponse struct {
	Name string `json:"name"`
	Nis  string `json:"nis"`
}
