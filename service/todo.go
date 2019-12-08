package service

import (
	"github.com/luiz-mai/go-api-boilerplate/data"
	"github.com/luiz-mai/go-api-boilerplate/model"
)

type ToDoService struct {
	data *data.Conn
}

func NewToDoService(db *data.Conn) ToDoService {
	return ToDoService{
		data: db,
	}
}

func (s ToDoService) GetAll() ([]model.ToDo, error) {
	return s.data.ToDo().GetAll()
}

func (s ToDoService) Create(todo model.ToDo) error {
	return s.data.ToDo().Create(todo)
}
