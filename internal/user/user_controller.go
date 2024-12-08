package user

import (
	"net/http"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/Lunarisnia/inventory-manager/internal/auth"
	"github.com/Lunarisnia/inventory-manager/internal/user/usermodels"
	"github.com/gin-gonic/gin"
)

// TODO: EASY: Add register user endpoint
// TODO: EASY: Add Change Password Endpoint
// TODO: MEDIUM: Add Login endpoint
// TODO: MEDIUM: Add Borrow item endpoint
// TODO: HARD: Add Return item endpoint with admin barcode validation via config
// TODO: HARD: Auth middleware

type UserController struct {
	repository   *repo.Queries
	tokenManager auth.TokenManager
}

func NewUserController(r *gin.RouterGroup, repository *repo.Queries) *UserController {
	group := r.Group("/user")

	ctl := UserController{
		repository: repository,
	}
	group.GET("/ping", ctl.Ping)
	group.GET("/pong", ctl.Pong)
	return &ctl
}

func (u *UserController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}

func (u *UserController) Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"bizz":  "buzz",
		"bizz1": "buzz2",
	})
}

func (u *UserController) Login(c *gin.Context) {
	userCredential := usermodels.UserLoginCredential{}
	if err := c.BindJSON(&userCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing message.",
		})
		return
	}

	user, err := u.repository.GetUserByNIS(c.Request.Context(), userCredential.NIS)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	token, err := u.tokenManager.Generate(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
