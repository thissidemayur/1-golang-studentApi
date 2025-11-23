package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thissidemayur/1-golang-studentsApi/internal/config"
	"github.com/thissidemayur/1-golang-studentsApi/internal/types"
)

type Sqlite struct{
	Db *sql.DB
}


// initalize db
func New(cfg *config.Config) (*Sqlite,error) {
	// pass driver
	db,err:=sql.Open("sqlite3",cfg.StoragePath)
	if err != nil {
		log.Fatalf("database connection error: %s\n",err.Error())
		return nil,err
	}

	// initalize table 
	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS student(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	rollNo INTEGER
	)`)

	if err != nil {
		return nil,err
	}

	return &Sqlite{
		Db: db,
	},nil
}



// implement storage interface methods here
// CreateStudent method
// //// (S* sqlite) -> pointer receiver to Sqlite struct
func (s *Sqlite)    CreateStudent(name string, email string, rollNo int) (int64,error) { 
	// prepare statement
	stmt,err:=s.Db.Prepare("INSERT INTO student (name, email,rollNo) VALUES(?,?,?)")

	if err != nil{
		fmt.Println("Preparing statmeent error: ",err)
		return 0,err
	}
	defer stmt.Close()
	// ecxecute query
	result,err :=stmt.Exec(name,email,	rollNo) 
	if err != nil {
		return 0,err
	}
	lastId,err := result.LastInsertId()
	 if err != nil {
		return 0,err
	 }
	return lastId,nil

}

func (s *Sqlite) GetStudentById(id int64) (types.Student,error) {

	// prepare query
	stmt,err := s.Db.Prepare("SELECT * FROM student where id = ? LIMIT 1"  ) 
	if err !=nil {
		return types.Student{} ,err
	}
	defer stmt.Close()

	var student types.Student

	// execute query
	err = stmt.QueryRow(id).Scan(&student.ID,&student.Name,&student.Email,&student.RollNo)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student with id %d not found",id)
		}
		return types.Student{}, err
	}

	return student,nil
}


func (s *Sqlite) GetStudentList() ([]types.Student,error) {

	// prepare query
	stmt,err:=s.Db.Prepare("SELECT * from student")
	if err != nil {
		return nil,err
	}
	defer stmt.Close()

	// execute query
	rows,err:=stmt.Query()
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	var students [] types.Student

	for rows.Next() {
		var student types.Student
		err:=rows.Scan(&student.ID, &student.Name,&student.Email,&student.RollNo)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}