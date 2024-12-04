package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: EASY: Add register user endpoint
// TODO: EASY: Add Change Password Endpoint
// TODO: MEDIUM: Add Login endpoint
// TODO: MEDIUM: Add Borrow item endpoint
// TODO: HARD: Add Return item endpoint with admin barcode validation via config
// TODO: HARD: Auth middleware

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
