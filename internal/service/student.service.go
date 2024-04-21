package service

import (
	"context"
	"simple-go-app/internal/model"
	"simple-go-app/internal/repository"
)

type StudentService interface {
	GetAll(ctx context.Context) (*[]model.Student, error)
	GetById(ctx context.Context, id int) (*model.Student, error)
}

type StudentServ struct {
	SR repository.StudentRepository
}

func (s *StudentServ) GetAll(ctx context.Context) (*[]model.Student, error) {
	students, err := s.SR.GetAllStudents(ctx)

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *StudentServ) GetById(ctx context.Context, id int) (*model.Student, error) {
	student, err := s.SR.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return student, nil
}
