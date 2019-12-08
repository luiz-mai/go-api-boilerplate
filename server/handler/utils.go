package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func responseOK(ctx echo.Context, i interface{}) error {
	return ctx.JSON(http.StatusOK, i)
}

func responseNoContent(ctx echo.Context) error {
	return ctx.JSON(http.StatusNoContent, nil)
}
