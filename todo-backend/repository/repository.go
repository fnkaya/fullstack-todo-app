package repository

import (
	"backend/model"
)

type IRepository interface {
	Insert(todo model.Todo) error
	GetAll() (model.Todos, error)
}

type database struct {
	todos model.Todos
}

func (db database) GetAll() (model.Todos, error) {
	return db.todos, nil
}

func (db *database) Insert(todo model.Todo) error {
	todo.Id = db.generateId()
	db.todos = append(db.todos, todo)
	return nil
}

func (db database) generateId() int {
	if len(db.todos) > 0 {
		return db.todos[len(db.todos)-1].Id + 1
	} else {
		return 1
	}
}

func NewRepository() IRepository {
	return &database{todos: model.Todos{
		model.Todo{
			Id: 1, Text: "todo-1", Done: true,
		},
	}}
}
