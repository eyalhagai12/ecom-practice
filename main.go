package main

import (
	"context"
	"ecom/server"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	app := echo.New()
	server := server.New(ctx, app)

	err := server.Run()
	if err != nil {
		log.Fatal("failed to run server: ", err.Error())
	}
}
