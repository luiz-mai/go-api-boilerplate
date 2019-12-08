package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/luiz-mai/go-api-boilerplate/config"
)

func Run(cfg config.Config) (err error) {
	e := echo.New()
	e.HideBanner = true

	err = e.Start(fmt.Sprintf(":%d", cfg.HTTP.Port))
	if err != nil {
		return err
	}

	return nil
}
