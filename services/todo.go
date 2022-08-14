package services

import (
	"github.com/bonifacio/todogo/models"
	"github.com/bonifacio/todogo/repositories"
)

type (
	TodoService interface {
		Create(todo models.Todo) *models.Todo
	}
	todoService struct {
		repository repositories.TodoRepository
	}
)

func (ts todoService) Create(todo models.Todo) *models.Todo {
	return ts.repository.Save(todo)
}

func NewTodoService(repository repositories.TodoRepository) TodoService {
	return todoService{
		repository,
	}
}
