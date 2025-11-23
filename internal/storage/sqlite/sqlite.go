package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thissidemayur/1-golang-studentsApi/internal/config"
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