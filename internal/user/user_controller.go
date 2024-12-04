package user

import (
	"net/http"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/gin-gonic/gin"
)

// TODO: EASY: Add register user endpoint
// TODO: EASY: Add Change Password Endpoint
// TODO: MEDIUM: Add Login endpoint
// TODO: MEDIUM: Add Borrow item endpoint
// TODO: HARD: Add Return item endpoint with admin barcode validation via config
// TODO: HARD: Auth middleware

type UserController struct {
	repository *repo.Queries
}

func NewUserController(r *gin.RouterGroup, repository *repo.Queries) *UserController {
	group := r.Group("/user")

	ctl := UserController{
		repository: repository,
	}
	group.GET("/ping", ctl.Ping)

	return &ctl
}

func (u *UserController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
