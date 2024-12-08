package auth

import (
	"context"
	"errors"
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

type tokenManagerImpl struct{}

func (t *tokenManagerImpl) Generate(ctx context.Context, user *repo.User) (string, error) {
	if user == nil {
		return "", errors.New("user can't be empty")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"nis": user.Nis,
		"exp": time.Now().Add(5 * time.Hour).Unix(),
	})
	return token.Raw, nil
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
		bearerToken := c.Request.Header["Authorization"][0]
		ctx, err := validate(c.Request.Context(), bearerToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// func Logger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		t := time.Now()
//
// 		// Set example variable
// 		c.Set("example", "12345")
//
// 		// before request
//
// 		c.Next()
//
// 		// after request
// 		latency := time.Since(t)
// 		log.Print(latency)
//
// 		// access the status we are sending
// 		status := c.Writer.Status()
// 		log.Println(status)
// 	}
// }