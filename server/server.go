package server

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/luiz-mai/go-api-boilerplate/config"
	"github.com/luiz-mai/go-api-boilerplate/server/handler"
	"github.com/luiz-mai/go-api-boilerplate/service"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
)

func Run(
	cfg config.Config,
	todoService service.ToDoService,
) {
	e := echo.New()
	e.HideBanner = true

	e.GET("/health", handler.HandleHealthCheck())

	e.GET("/todos", handler.HandleGetToDos(todoService))
	e.GET("/todos/:id", handler.HandleGetToDo(todoService))
	e.POST("/todos", handler.HandleCreateToDo(todoService))
	e.PUT("/todos/:id", handler.HandleUpdateToDo(todoService))
	e.DELETE("/todos/:id", handler.HandleDeleteToDo(todoService))

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.HTTP.Port)); err != nil {
			log.Warn("shutting down server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
