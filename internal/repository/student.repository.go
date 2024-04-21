package repository

import (
	"context"
	"database/sql"
	"simple-go-app/internal/model"
)

type StudentRepository interface {
	GetAllStudents(ctx context.Context) (*[]model.Student, error)
	GetById(ctx context.Context, id int) (*model.Student, error)
	// Create(ctx context.Context, student *model.Student) (int, error)
	// Update(ctx context.Context, student *model.Student, id int) (*model.Student, error)
	// Delete(ctx context.Context, id int) (int, error)
}

type StudentRepo struct {
	DB *sql.DB
}

func (s *StudentRepo) GetAllStudents(ctx context.Context) (*[]model.Student, error) {
	var students []model.Student
	rows, err := s.DB.QueryContext(ctx, "select * from student")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var student model.Student
		if err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Age); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}

func (s *StudentRepo) GetById(ctx context.Context, id int) (*model.Student, error) {
	var student model.Student
	err := s.DB.QueryRowContext(ctx, "select * from student where id=$1", id).Scan(
		&student.ID, &student.Name, &student.Email, &student.Age)

	if err != nil {
		return nil, err
	}

	return &student, nil
}
