package service

import (
	"context"
	"simple-go-app/internal/model"
	"simple-go-app/internal/repository"
)

type StudentService interface {
	GetAll(ctx context.Context) (*[]model.Student, error)
	GetById(ctx context.Context, id int) (*model.Student, error)
	Create(ctx context.Context, student *model.Student) (int, error)
	Update(ctx context.Context, student *model.Student, id int) (*model.Student, error)
	Delete(ctx context.Context, id int) (int, error)
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

func (s *StudentServ) Create(ctx context.Context, student *model.Student) (int, error) {
	id, err := s.SR.Create(ctx, student)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *StudentServ) Update(ctx context.Context, student *model.Student, id int) (*model.Student, error) {
	updatedStudent, err := s.SR.Update(ctx, student, id)

	if err != nil {
		return nil, err
	}

	return updatedStudent, nil
}

func (s *StudentServ) Delete(ctx context.Context, id int) (int, error) {
	affectedRows, err := s.SR.Delete(ctx, id)

	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}
