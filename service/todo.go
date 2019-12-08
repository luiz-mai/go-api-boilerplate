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

func (s ToDoService) GetByID(id int64) (model.ToDo, error) {
	return s.data.ToDo().GetByID(id)
}

func (s ToDoService) Create(todo model.ToDo) error {
	return s.data.ToDo().Create(todo)
}

func (s ToDoService) Update(todo model.ToDo) error {
	return s.data.ToDo().Update(todo)
}

func (s ToDoService) Delete(id int64) error {
	return s.data.ToDo().Delete(id)
}
