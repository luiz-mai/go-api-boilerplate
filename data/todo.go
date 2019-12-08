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

func (r todoRepo) GetByID(id int64) (model.ToDo, error) {
	const query = `
		SELECT
			todo_id
			, title
			FROM tab_todo
			WHERE todo_id = ?
		;
	`

	todo, err := r.parse(r.conn.QueryRow(query, id))
	if err != nil {
		return model.ToDo{}, err
	}

	return todo, nil
}

func (r todoRepo) Create(todo model.ToDo) error {
	const query = `
		INSERT INTO tab_todo
			( title
			)
			VALUES
				( ?
				)
		;
	`

	_, err := r.conn.Exec(query, todo.Title)
	if err != nil {
		return err
	}

	return nil
}

func (r todoRepo) Update(todo model.ToDo) error {
	const query = `
		UPDATE tab_todo
			SET title = ?
			WHERE todo_id = ?
		;
	`

	_, err := r.conn.Exec(query, todo.Title, todo.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r todoRepo) Delete(todo model.ToDo) error {
	const query = `
		DELETE FROM tab_todo
			WHERE todo_id = ?
		;
	`

	_, err := r.conn.Exec(query, todo.ID)
	if err != nil {
		return err
	}

	return nil
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
