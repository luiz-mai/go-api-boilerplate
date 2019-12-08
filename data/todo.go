package data

import (
	"database/sql"

	"github.com/luiz-mai/go-api-boilerplate/model"
)

type todoRepo struct {
	conn executor
}

func newTodoRepo(conn executor) todoRepo {
	return todoRepo{
		conn: conn,
	}
}

func (r todoRepo) GetAll() ([]model.ToDo, error) {
	const query = `
		SELECT
			todo_id
			, title
			FROM tab_todo
		;
	`

	todos, err := r.parseList(r.conn.Query(query))
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r todoRepo) parseList(rows *sql.Rows, err error) ([]model.ToDo, error) {
	if err != nil {
		return nil, err
	}

	todos := make([]model.ToDo, 0)
	for rows.Next() {

		todo, err := r.parse(rows)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r todoRepo) parse(s scanner) (model.ToDo, error) {

	var todo model.ToDo
	err := s.Scan(
		&todo.ID,
		&todo.Title,
	)
	if err != nil {
		return model.ToDo{}, err
	}

	return todo, nil
}
