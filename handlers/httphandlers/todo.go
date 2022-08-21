package httphandlers

import (
	"net/http"

	"github.com/bonifacio/todogo/models"
	"github.com/gin-gonic/gin"
)

type (
	TodoService interface {
		Create(todo models.Todo) *models.Todo
	}
	TodoHandler struct {
		service TodoService
	}
)

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

func (th TodoHandler) Upload(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	file, _ := fileHeader.Open()
	b := make([]byte, fileHeader.Size)
	_, _ = file.Read(b)
	c.JSON(http.StatusOK, gin.H{
		"content": string(b),
	})
}

func ConfigureTodoHandler(engine *gin.Engine, handler *TodoHandler) {
	engine.POST("/todo", handler.CreateTodo)
	engine.POST("/todo/:id/file", handler.Upload)
}

func NewTodoHandler(service TodoService) *TodoHandler {
	return &TodoHandler{service}
}
