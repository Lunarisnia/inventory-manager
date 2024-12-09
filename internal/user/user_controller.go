package user

import (
	"net/http"
	"time"

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

func NewUserController(r *gin.RouterGroup, repository *repo.Queries, tokenManager auth.TokenManager) *UserController {
	group := r.Group("/user")

	ctl := UserController{
		repository:   repository,
		tokenManager: tokenManager,
	}
	group.GET("/ping", auth.Authorized(), ctl.Ping)
	group.GET("/pong", ctl.Pong)
	group.POST("/login", ctl.Login)
	group.POST("/register", ctl.Register)
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

func (u *UserController) Register(c *gin.Context) {
	newUser := usermodels.RegisterUser{}
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request you dingus.",
		})
		return
	}

	user, err := u.repository.CreateUser(c.Request.Context(), repo.CreateUserParams{
		Name:      newUser.Name,
		Nis:       newUser.NIS,
		Password:  newUser.Password,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
