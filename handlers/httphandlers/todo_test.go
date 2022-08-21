package httphandlers_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bonifacio/todogo/handlers/httphandlers"
	"github.com/bonifacio/todogo/handlers/httphandlers/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGiven_When_Then(t *testing.T) {

	service := new(mocks.TodoService)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "file.txt")
	assert.NoError(t, err)
	file, err := os.Open("mocks/file.txt")
	assert.NoError(t, err)

	_, err = io.Copy(part, file)
	assert.NoError(t, err)
	assert.NoError(t, writer.Close())

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/todo/1/file", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())

	handle := httphandlers.NewTodoHandler(service)

	router := setupRouter(handle)
	router.ServeHTTP(w, r)

	responseBody, _ := io.ReadAll(w.Result().Body)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, `{"content":"mock"}`, string(responseBody))
}

func setupRouter(h *httphandlers.TodoHandler) *gin.Engine {
	g := gin.Default()
	httphandlers.ConfigureTodoHandler(g, h)
	return g
}
