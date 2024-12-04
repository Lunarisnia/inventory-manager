package core

import (
	"context"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/Lunarisnia/inventory-manager/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type App struct {
	ctx    context.Context
	engine *gin.Engine
	db     *pgx.Conn
}

func NewApp(ctx context.Context) *App {
	e := gin.Default()
	r := e.Group("/v1")

	conn, err := db.Connect(ctx, "postgres://root:password@localhost:5432/inventory-manager")
	if err != nil {
		panic(err)
	}
	repository := repo.New(conn)
	InitializeRouter(r, repository)
	return &App{
		ctx:    ctx,
		engine: e,
		db:     conn,
	}
}

func (a *App) Run() {
	defer a.db.Close(a.ctx)
	a.engine.Run(":3009") // listen and serve on 0.0.0.0:8080
}
