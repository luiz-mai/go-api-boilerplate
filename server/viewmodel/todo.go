package viewmodel

import "github.com/luiz-mai/go-api-boilerplate/model"

type GetToDoResponse struct {
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func ParseGetToDosResponse(todos []model.ToDo) []GetToDoResponse {
	response := make([]GetToDoResponse, len(todos))
	for i := range todos {
		response[i] = GetToDoResponse{
			ID:    todos[i].ID,
			Title: todos[i].Title,
		}
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
