package repositories

import (
	"log"

	"github.com/bonifacio/todogo/models"
)

type (
	TodoRepository interface {
		Save(todo models.Todo) *models.Todo
	}
	todoRepository struct {
		lista []models.Todo
	}
)

func (tr todoRepository) Save(todo models.Todo) *models.Todo {
	tr.lista = append(tr.lista, todo)
	log.Default().Println(tr.lista)
	return &todo
}

func NewTodoRepository() TodoRepository {
	return todoRepository{
		lista: make([]models.Todo, 0),
	}
}
