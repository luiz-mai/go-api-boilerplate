package handler

import (
	"github.com/labstack/echo"
	"github.com/luiz-mai/go-api-boilerplate/service"
	"github.com/pkg/errors"
)

func HandleGetToDos(todoService service.ToDoService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		todos, err := todoService.GetAll()
		if err != nil {
			return errors.New("Error while getting list of ToDos")
		}

		return ResponseOK(ctx, todos)
	}
}
