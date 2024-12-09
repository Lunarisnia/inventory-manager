package auth

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	UserID int32  `json:"user_id"`
	NIS    string `json:"nis"`
	Exp    int64  `json:"exp"`
	jwt.RegisteredClaims
}
