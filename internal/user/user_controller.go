package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController(r *gin.RouterGroup) *UserController {
	group := r.Group("/user")

	ctl := UserController{}
	group.GET("/ping", ctl.Ping)

	return &ctl
}

func (u *UserController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
