package service

import (
	"backend/model"
	"backend/repository"
	"errors"
)

var (
	EmptyTextError = errors.New("text can not be empty")
)

type IService interface {
	Insert(todo model.Todo) error
	GetAll() (model.Todos, error)
}

type Service struct {
	repository repository.IRepository
}

func (s *Service) Insert(todo model.Todo) error {
	if todo.Text == "" {
		return EmptyTextError
	}

	return s.repository.Insert(todo)
}

func (s Service) GetAll() (model.Todos, error) {
	return s.repository.GetAll()
}

func NewService(repo repository.IRepository) IService {
	return &Service{repository: repo}
}
