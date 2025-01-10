package handlers

import "github.com/labstack/echo/v4"

type handlerFunc[Req any, Res any] func(echo.Context, Req) (Res, error)

func FromFunc[Req any, Res any](handler handlerFunc[Req, Res], successCode int) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req Req
		if err := c.Bind(req); err != nil {
			return err
		}

		res, err := handler(c, req)
		if err != nil {
			return err
		}

		return c.JSON(successCode, res)
	}
}
