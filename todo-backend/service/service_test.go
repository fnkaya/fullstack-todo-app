package service

import (
	"backend/mocks"
	"backend/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {

	mockRepository := mocks.NewMockIRepository(gomock.NewController(t))

	t.Run("should return error if text is empty", func(t *testing.T) {
		mockTodo := model.Todo{}
		service := NewService(mockRepository)
		err := service.Insert(mockTodo)

		assert.Equal(t, EmptyTextError, err)
	})

	t.Run("should delegate repository insert func", func(t *testing.T) {
		mockTodo := model.Todo{Text: "new todo"}
		mockRepository.EXPECT().
			Insert(mockTodo).
			Return(nil).
			Times(1)

		service := NewService(mockRepository)
		err := service.Insert(mockTodo)

		assert.Nil(t, err)
	})

	t.Run("should delegate repository getAll func", func(t *testing.T) {
		mockTodos := model.Todos{
			model.Todo{Id: 1, Text: "todo-1", Done: false},
			model.Todo{Id: 2, Text: "todo-2", Done: false},
			model.Todo{Id: 3, Text: "todo-3", Done: true},
		}
		mockRepository.EXPECT().
			GetAll().
			Return(mockTodos, nil).
			Times(1)

		service := NewService(mockRepository)
		todos, err := service.GetAll()

		assert.ObjectsAreEqualValues(mockTodos, todos)
		assert.Nil(t, err)
	})
}
