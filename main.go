package main

import (
	"context"

	"github.com/Lunarisnia/inventory-manager/internal/core"
)

func main() {
	app := core.NewApp(context.Background())
	app.Run()
}
