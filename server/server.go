package server

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Server struct {
	app *echo.Echo
	env Env
}

func New(ctx context.Context, app *echo.Echo) Server {
	env := NewEnv(ctx)

	return Server{
		app: app,
		env: env,
	}
}

func (s Server) Run() error {
	return s.app.Start(":8080")
}
