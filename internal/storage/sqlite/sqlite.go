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



