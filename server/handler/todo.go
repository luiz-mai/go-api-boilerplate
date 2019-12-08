package handler

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/luiz-mai/go-api-boilerplate/server/viewmodel"
	"github.com/luiz-mai/go-api-boilerplate/service"
	"github.com/pkg/errors"
)

func HandleGetToDos(todoService service.ToDoService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		todos, err := todoService.GetAll()
		if err != nil {
			return errors.New("Error while getting list of ToDos")
		}

		vm := viewmodel.ParseGetToDosResponse(todos)

		return responseOK(ctx, vm)
	}
}

func HandleGetToDo(todoService service.ToDoService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			fmt.Println(err)
			return errors.New("Invalid id")
		}

		todo, err := todoService.GetByID(id)
		if err != nil {
			return errors.New("Error while getting list of ToDos")
		}

		vm := viewmodel.ParseGetToDoResponse(todo)

		return responseOK(ctx, vm)
	}
}

func HandleCreateToDo(todoService service.ToDoService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		var vm viewmodel.CreateToDoRequest
		err := ctx.Bind(&vm)
		if err != nil {
			fmt.Println(err)
			return errors.New("Error parsing request body")
		}

		todo := vm.Parse()
		err = todoService.Create(todo)
		if err != nil {
			fmt.Println(err)
			return errors.New("Error creating a new ToDo")
		}

		return responseNoContent(ctx)
	}
}

func HandleUpdateToDo(todoService service.ToDoService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			fmt.Println(err)
			return errors.New("Invalid id")
		}

		var vm viewmodel.UpdateToDoRequest
		err = ctx.Bind(&vm)
		if err != nil {
			fmt.Println(err)
			return errors.New("Error parsing request body")
		}

		todo := vm.Parse()
		todo.ID = id

		err = todoService.Update(todo)
		if err != nil {
			fmt.Println(err)
			return errors.New("Error creating a new ToDo")
		}

		return responseNoContent(ctx)
	}
}

func HandleDeleteToDo(todoService service.ToDoService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			fmt.Println(err)
			return errors.New("Invalid id")
		}

		err = todoService.Delete(id)
		if err != nil {
			fmt.Println(err)
			return errors.New("Error creating a new ToDo")
		}

		return responseNoContent(ctx)
	}
}
