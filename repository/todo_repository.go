package repository

import (
	"github.com/alekseinovikov/todo/domain"
)

type TodoRepository struct {
	storage map[string]domain.Todo
}

func NewTodoRepository() *TodoRepository {
	storage := make(map[string]domain.Todo)
	return &TodoRepository{storage: storage}
}

func (t TodoRepository) GetById(id string) (error, domain.Todo) {
	if todo, ok := t.storage[id]; ok {
		return nil, todo
	}

	return domain.ErrTodoNotFound, domain.Todo{}
}

func (t TodoRepository) Create(todo domain.Todo) (error, domain.Todo) {
	t.storage[todo.Id] = todo
	return nil, todo
}
