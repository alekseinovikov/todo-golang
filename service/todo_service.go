package service

import (
	"github.com/alekseinovikov/todo/domain"
	uuid "github.com/google/uuid"
)

type TodoRepository interface {
	GetById(id string) (error, domain.Todo)
	Create(todo domain.Todo) (error, domain.Todo)
	Delete(id string) error
}

type TodoService struct {
	repository TodoRepository
}

func NewTodoService(repository TodoRepository) *TodoService {
	return &TodoService{repository: repository}
}

func (t TodoService) GetById(id string) (error, domain.Todo) {
	return t.repository.GetById(id)
}

func (t TodoService) Create(todo domain.Todo) (error, domain.Todo) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err, domain.Todo{}
	}

	todo.Id = newUUID.String()
	return t.repository.Create(todo)
}

func (t TodoService) Delete(id string) error {
	err, _ := t.repository.GetById(id)
	if err != nil {
		return err
	}

	return t.repository.Delete(id)
}
