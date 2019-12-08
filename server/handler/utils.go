package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func ResponseOK(ctx echo.Context, i interface{}) error {
	return ctx.JSON(http.StatusOK, i)
}
