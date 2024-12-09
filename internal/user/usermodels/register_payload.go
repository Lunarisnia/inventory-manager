package usermodels

type RegisterUser struct {
	Name     string `json:"name"`
	NIS      string `json:"nis"`
	Password string `json:"password"`
}
