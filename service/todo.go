package service

import "github.com/luiz-mai/go-api-boilerplate/model"

type ToDoService struct{}

func NewToDoService() ToDoService {
	return ToDoService{}
}

func (s ToDoService) GetAll() ([]model.ToDo, error) {
	return []model.ToDo{}, nil
}
