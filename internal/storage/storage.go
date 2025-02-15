package storage

import "github.com/ak-er/golang-playground/internal/schema"

type Storager interface {
	GetStudentList() ([]schema.Student, error)
	CreateStudent(name string, course string, city string, age int) (int64, error)
	GetStudentById(id int64) (schema.Student, error)
	DeleteStudent(id int64) (int64, error)
}
