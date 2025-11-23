package storage

import "github.com/thissidemayur/1-golang-studentsApi/internal/types"


type Storage interface {
	// define storage methods here
	CreateStudent(name string, email string, rollNo int) (int64,error)
	GetStudentById(id int64) (types.Student,error)
}

