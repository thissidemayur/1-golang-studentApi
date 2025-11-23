package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

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

// GetStudentById method
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

// GetStudentList method
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

// DeleteStudentById method
func (s *Sqlite) DeleteStudentById(id int64) error {

	// prepare statement
	stmt, err := s.Db.Prepare("DELETE FROM student WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute statement
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateStudentById method
func (s *Sqlite) UpdateStudentById(id int64, dto types.UpdateStudent) (types.Student,error){
	query := "UPDATE student SET "
	args := []interface{}{}
	updates := []string{}

	if dto.Name != nil {
		updates = append(updates,"name = ? ")
		args = append(args,*dto.Name)

	}
	if dto.Email != nil {
		updates = append(updates,"email = ? ")
		args = append(args,*dto.Email)
	}
	if dto.RollNo != nil {
		updates = append(updates,"rollNo = ? ")
		args = append(args,*dto.RollNo)
	}

	if len(updates) == 0 {
		return types.Student{} , errors.New("no fields to update")
	}

	// join updates
	query+=strings.Join(updates, ", ") + " WHERE id = ?"
	args = append(args, id)

	// prepare statement
	data,err:=s.Db.Exec(query,args...)
	if err != nil {
		return types.Student{}, err
	}
	afftectedId,err :=data.RowsAffected()
	if err != nil {
		return types.Student{}, err
	}
	if afftectedId == 0 {
		return types.Student{}, fmt.Errorf("student with id %d not found",id)
	}


	student, err := s.GetStudentById(id) 
	if err != nil {
		return types.Student{}, err
	}
	return student, nil
}