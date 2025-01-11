package main

import (
	"context"
	"ecom/server"
	"ecom/store"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx := context.Background()

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	server := server.New(ctx, app)

	server.RegsiterHandlers(
		store.RegisterHandlers,
	)

	err := server.Run()
	if err != nil {
		log.Fatal("failed to run server: ", err.Error())
	}
}
