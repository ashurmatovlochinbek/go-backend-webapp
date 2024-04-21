package repository

import (
	"context"
	"database/sql"
	"log"
	"simple-go-app/internal/model"
)

type StudentRepository interface {
	GetAllStudents(ctx context.Context) (*[]model.Student, error)
	GetById(ctx context.Context, id int) (*model.Student, error)
	Create(ctx context.Context, student *model.Student) (int, error)
	Update(ctx context.Context, student *model.Student, id int) (*model.Student, error)
	// Delete(ctx context.Context, id int) (int, error)
}

type StudentRepo struct {
	DB *sql.DB
}

func (s *StudentRepo) GetAllStudents(ctx context.Context) (*[]model.Student, error) {
	var students []model.Student
	rows, err := s.DB.QueryContext(ctx, "select * from student")

	if err != nil {
		log.Printf("Could not get all students: %v", err)
		return nil, err
	}

	for rows.Next() {
		var student model.Student
		if err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Age); err != nil {
			log.Printf("Could not get student by scan method: %v", err)
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
		log.Printf("Could not get student by id: %v", err)
		return nil, err
	}

	return &student, nil
}

func (s *StudentRepo) Create(ctx context.Context, student *model.Student) (int, error) {
	var id int
	err := s.DB.QueryRowContext(ctx, "insert into student(name, email, age) values($1, $2, $3) returning id",
		student.Name,
		student.Email,
		student.Age,
	).Scan(&id)

	if err != nil {
		log.Printf("Could not create student: %v", err)
		return 0, err
	}

	return id, nil
}

func (s *StudentRepo) Update(ctx context.Context, student *model.Student, id int) (*model.Student, error) {
	var oldStudent model.Student
	err := s.DB.QueryRowContext(ctx, "select * from student where id=$1", id).Scan(
		&oldStudent.ID,
		&oldStudent.Name,
		&oldStudent.Email,
		&oldStudent.Age,
	)

	if err != nil {
		return nil, err
	}

	if student.Name != "" {
		oldStudent.Name = student.Name
	}

	if student.Email != "" {
		oldStudent.Email = student.Email
	}

	if student.Age != 0 {
		oldStudent.Age = student.Age
	}

	_, err = s.DB.ExecContext(ctx, "update student set name=$1, email=$2, age=$3 where id=$4",
		oldStudent.Name, oldStudent.Email, oldStudent.Age, oldStudent.ID,
	)

	if err != nil {
		return nil, err
	}

	return &oldStudent, nil
}
