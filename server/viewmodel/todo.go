package viewmodel

import "github.com/luiz-mai/go-api-boilerplate/model"

type GetToDoResponse struct {
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func ParseGetToDoResponse(todo model.ToDo) GetToDoResponse {
	return GetToDoResponse{
		ID:    todo.ID,
		Title: todo.Title,
	}
}

func ParseGetToDosResponse(todos []model.ToDo) []GetToDoResponse {
	response := make([]GetToDoResponse, len(todos))
	for i := range todos {
		response[i] = ParseGetToDoResponse(todos[i])
	}

	return response
}

type CreateToDoRequest struct {
	Title string `json:"title"`
}

func (vm CreateToDoRequest) Parse() model.ToDo {
	return model.ToDo{
		Title: vm.Title,
	}
}

type UpdateToDoRequest struct {
	Title string `json:"title"`
}

func (vm UpdateToDoRequest) Parse() model.ToDo {
	return model.ToDo{
		Title: vm.Title,
	}
}
