package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type FuncHandler[Request any, Response any] func(echo.Context, Env, Request) (Response, error)

func HandlerFromFunc[Request any, Response any](env Env, handler FuncHandler[Request, Response], successCode int) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request Request
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request: "+err.Error())
		}

		if err := validator.New().Struct(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to validate request: "+err.Error())
		}

		response, err := handler(c, env, request)
		if err != nil {
			return err
		}

		return c.JSON(successCode, response)
	}
}

type NoRequestFuncHandler[Response any] func(echo.Context, Env) (Response, error)

func HandlerNoRequestFromFunc[Response any](env Env, handler NoRequestFuncHandler[Response], successCode int) echo.HandlerFunc {
	return func(c echo.Context) error {
		response, err := handler(c, env)
		if err != nil {
			return err
		}

		return c.JSON(successCode, response)
	}
}
