package auth

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	ID  int32  `json:"id"`
	NIS string `json:"nis"`
	Exp int64  `json:"exp"`
	jwt.RegisteredClaims
}
