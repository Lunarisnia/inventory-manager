package core

import (
	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/Lunarisnia/inventory-manager/internal/auth"
	"github.com/Lunarisnia/inventory-manager/internal/borrowlist"
	"github.com/Lunarisnia/inventory-manager/internal/item"
	"github.com/Lunarisnia/inventory-manager/internal/user"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	UserController       *user.UserController
	BorrowListController *borrowlist.BorrowListController
	ItemController       *item.ItemController
}

func InitializeRouter(r *gin.RouterGroup, repository *repo.Queries) *RouterGroup {
	tokenManager := auth.NewTokenManager()
	return &RouterGroup{
		UserController:       user.NewUserController(r, repository, tokenManager),
		BorrowListController: borrowlist.NewBorrowListController(r, repository),
		ItemController:       item.NewItemController(r, repository),
	}
}
