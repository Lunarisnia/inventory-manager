package core

import (
	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/Lunarisnia/inventory-manager/internal/user"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	UserController *user.UserController
}

func InitializeRouter(r *gin.RouterGroup, repository *repo.Queries) *RouterGroup {
	return &RouterGroup{
		UserController: user.NewUserController(r),
	}
}
