package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/ak-er/golang-playground/internal/config"
	"github.com/ak-er/golang-playground/internal/schema"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func (s *Sqlite) GetStudentList() ([]schema.Student, error) {
	stmt, err := s.Db.Prepare(`SELECT * FROM students`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []schema.Student
	for rows.Next() {
		var student schema.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Course, &student.City, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *Sqlite) CreateStudent(name string, course string, city string, age int) (int64, error) {
	stmt, err := s.Db.Prepare(`INSERT INTO students (name, course, city, age) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, course, city, age)
	if err != nil {
		return 0, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertedId, nil
}

func (s *Sqlite) GetStudentById(id int64) (schema.Student, error) {
	student := schema.Student{}
	stmt, err := s.Db.Prepare(`SELECT * FROM students WHERE id=? LIMIT 1`)
	if err != nil {
		return student, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Course, &student.City, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return schema.Student{}, fmt.Errorf("Invalid credentials")
		}
		return schema.Student{}, err
	}
	return student, nil
}

func (s *Sqlite) DeleteStudent(id int64) (int64, error) {
	stmt, err := s.Db.Prepare(`DELETE FROM students WHERE id=?`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, nil
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.Database)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name string,
	course string,
	city string,
	age INTEGER
	)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{Db: db}, nil
}
