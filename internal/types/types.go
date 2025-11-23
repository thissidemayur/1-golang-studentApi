package types

type Student struct {
	ID    int64
	Name  string `validate:"required,min=5,max=100"`
	Email string `validate:"required,email"`
	RollNo int `validate:"required"`
}

type UpdateStudent struct {
	Name *string `json:"name"`
	Email *string `json:"email" validate:"email"` 
	RollNo *int `json:"roll_no"`
}