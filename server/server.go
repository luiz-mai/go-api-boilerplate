package server

import (
	"fmt"

	"github.com/luiz-mai/go-api-boilerplate/config"
	"github.com/luiz-mai/go-api-boilerplate/server/handler"
	"github.com/luiz-mai/go-api-boilerplate/service"

	"github.com/labstack/echo"
)

func Run(
	cfg config.Config,
	todoService service.ToDoService,
) (err error) {
	e := echo.New()
	e.HideBanner = true

	e.GET("/health", handler.HandleHealthCheck())

	e.GET("/todos", handler.HandleGetToDos(todoService))
	e.GET("/todos/:id", handler.HandleGetToDo(todoService))
	e.POST("/todos", handler.HandleCreateToDo(todoService))
	e.PUT("/todos/:id", handler.HandleUpdateToDo(todoService))
	e.DELETE("/todos/:id", handler.HandleDeleteToDo(todoService))

	err = e.Start(fmt.Sprintf(":%d", cfg.HTTP.Port))
	if err != nil {
		return err
	}

	return nil
}
