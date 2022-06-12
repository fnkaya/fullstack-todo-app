package handler

import (
	"backend/mocks"
	"backend/model"
	"backend/service"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	todosUrl := "/api/todos"
	mockService := mocks.NewMockIService(gomock.NewController(t))

	t.Run("insert method test", func(t *testing.T) {
		t.Run("should return status NotImplemented", func(t *testing.T) {
			handler := NewHandler(mockService)
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/1", todosUrl), http.NoBody)
			res := httptest.NewRecorder()

			handler.HandleRequest(res, req)
			assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
		})

		t.Run("should return status CREATED if body is valid", func(t *testing.T) {
			mockTodo := model.Todo{Text: "nex todo"}
			mockService.EXPECT().
				Insert(mockTodo).
				Return(nil).
				Times(1)

			handler := NewHandler(mockService)
			data, _ := json.Marshal(mockTodo)
			req := httptest.NewRequest(http.MethodPost, todosUrl, bytes.NewReader(data))
			res := httptest.NewRecorder()

			handler.insert(res, req)
			assert.Equal(t, http.StatusCreated, res.Result().StatusCode)
		})

		t.Run("should return status InternalServiceError if body is not valid", func(t *testing.T) {
			handler := NewHandler(mockService)
			req := httptest.NewRequest(http.MethodPost, todosUrl, http.NoBody)
			res := httptest.NewRecorder()

			handler.insert(res, req)
			assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
		})

		t.Run("should return status InternalServiceError when service return an error", func(t *testing.T) {
			mockTodo := model.Todo{Text: ""}
			mockService.EXPECT().
				Insert(mockTodo).
				Return(service.EmptyTextError).
				Times(1)

			handler := NewHandler(mockService)
			data, _ := json.Marshal(mockTodo)
			req := httptest.NewRequest(http.MethodPost, todosUrl, bytes.NewReader(data))
			res := httptest.NewRecorder()

			handler.insert(res, req)
			assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
		})
	})

	t.Run("getAll method tests", func(t *testing.T) {
		t.Run("should return status OK", func(t *testing.T) {
			mockTodos := model.Todos{
				model.Todo{Id: 1, Text: "todo-1", Done: true},
				model.Todo{Id: 1, Text: "todo-2", Done: false},
			}
			mockService.EXPECT().
				GetAll().
				Return(mockTodos, nil).
				Times(1)

			handler := NewHandler(mockService)
			req := httptest.NewRequest(http.MethodGet, todosUrl, http.NoBody)
			res := httptest.NewRecorder()

			handler.getAll(res, req)
			assert.Equal(t, http.StatusOK, res.Result().StatusCode)
			assert.Equal(t, "application/json; charset=UTF-8", res.Header().Get("Content-Type"))
			response := model.Todos{}
			json.Unmarshal(res.Body.Bytes(), &response)
			assert.Equal(t, mockTodos, response)
		})

		t.Run("should return status InternalServerError when service return an error", func(t *testing.T) {
			serviceError := errors.New("any service error")
			mockService.EXPECT().
				GetAll().
				Return(nil, serviceError).
				Times(1)

			handler := NewHandler(mockService)
			req := httptest.NewRequest(http.MethodGet, todosUrl, http.NoBody)
			res := httptest.NewRecorder()

			handler.getAll(res, req)
			assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
			assert.Equal(t, []byte(serviceError.Error()), res.Body.Bytes())
		})
	})
}
