package services_test

import (
	"testing"

	"github.com/bonifacio/todogo/models"
	"github.com/bonifacio/todogo/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type todoRepositoryMock struct {
	mock.Mock
}

func (tr todoRepositoryMock) Save(todo models.Todo) *models.Todo {
	args := tr.Called(todo)

	response := args.Get(0)
	if response == nil {
		return nil
	}
	return response.(*models.Todo)
}

func TestGivenAnTodo_WhenSave_ThenSave(t *testing.T) {

	repositoryMock := new(todoRepositoryMock)
	todo := models.Todo{Description: "Testar o código"}
	repositoryMock.On("Save", todo).Return(&todo)

	service := services.NewTodoService(repositoryMock)

	response := service.Create(todo)

	assert.Equal(t, todo, *response)
}

func TestGivenAnTodo_WhenSave_ThenError(t *testing.T) {

	repositoryMock := new(todoRepositoryMock)
	todo := models.Todo{Description: "Testar o código"}
	repositoryMock.On("Save", todo).Return(nil)

	service := services.NewTodoService(repositoryMock)

	response := service.Create(todo)

	assert.Nil(t, response)
}
