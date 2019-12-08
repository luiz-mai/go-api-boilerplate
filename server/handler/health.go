package handler

import (
	"github.com/labstack/echo"
)

func HandleHealthCheck() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		return responseOK(ctx, "OK")
	}
}
