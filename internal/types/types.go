package types

type Student struct {
	ID    int
	Name  string `validate:"required,min=5,max=100"`
	Email string `validate:"required,email"`
	RollNo int `validate:"required"`
}