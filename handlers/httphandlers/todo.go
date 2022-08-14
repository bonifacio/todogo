package httphandlers

import (
	"net/http"

	"github.com/bonifacio/todogo/models"
	"github.com/bonifacio/todogo/services"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service services.TodoService
}

func (th TodoHandler) CreateTodo(c *gin.Context) {

	var todo models.Todo

	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	th.service.Create(todo)
	c.JSON(http.StatusCreated, todo)
}

func ConfigureTodoHandler(engine *gin.Engine, handler *TodoHandler) {
	engine.POST("/todo", handler.CreateTodo)
}

func NewTodoHandler(service services.TodoService) *TodoHandler {
	return &TodoHandler{service}
}
