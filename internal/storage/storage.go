package storage


type Storage interface {
	// define storage methods here
	CreateStudent(name string, email string, rollNo int) (int64,error)
}
