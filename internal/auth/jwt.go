package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	Generate(ctx context.Context, user *repo.User) (string, error)
}

func NewTokenManager() TokenManager {
	return &tokenManagerImpl{}
}

type tokenManagerImpl struct{}

func (t *tokenManagerImpl) Generate(ctx context.Context, user *repo.User) (string, error) {
	if user == nil {
		return "", errors.New("user can't be empty")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"nis":     user.Nis,
		"exp":     time.Now().Add(5 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte("foobar"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validate(ctx context.Context, bearerToken string) (context.Context, error) {
	bearerToken = strings.ReplaceAll(bearerToken, "Bearer ", "")
	claim := JWTClaim{}
	token, err := jwt.ParseWithClaims(bearerToken, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, "jwt", token.Claims)
	return ctx, nil
}

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header["Authorization"]
		if len(bearerToken) <= 0 {
			fmt.Println("THIS ONE")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		ctx, err := validate(c.Request.Context(), bearerToken[0])
		if err != nil {
			fmt.Println("This Other One")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
