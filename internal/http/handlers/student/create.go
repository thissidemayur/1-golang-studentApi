package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/thissidemayur/1-golang-studentsApi/internal/storage"
	"github.com/thissidemayur/1-golang-studentsApi/internal/types"
	"github.com/thissidemayur/1-golang-studentsApi/internal/utils/response"
)


func New(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// convert request data to struct
	var student types.Student

	// decode json data to struct
	 err:=json.NewDecoder(r.Body).Decode(&student)
	if errors.Is(err,io.EOF) {
		response.WriteJsonResponse(w,http.StatusBadRequest,response.GeneralError(fmt.Errorf("empty body")))
		return
	}

	if err != nil {
		response.WriteJsonResponse(w,http.StatusBadRequest,response.GeneralError(err))
		return
	}

	// request data validation
	if err := validator.New().Struct(student); err != nil {
		validateError := err.(validator.ValidationErrors)
		response.WriteJsonResponse(w,http.StatusBadRequest,response.ValidationError(validateError))
		return
	}

	// save info in db
	lastId,err :=storage.CreateStudent(student.Name,student.Email,student.RollNo)
	if err != nil {
		response.WriteJsonResponse(w,http.StatusInternalServerError,response.GeneralError(err))
		return
	}
	slog.Info("User created successfully. ",slog.String("userId",fmt.Sprintf("%d \n",lastId)))
	response.WriteJsonResponse(w,http.StatusCreated,map[string]string{
		"success":"true",
		"id":fmt.Sprint(lastId),
		"message":"Student created successfully",
	})

	}
	

}



/*
func New(){}  || func Create() {}  -> convention for handler function name
*/
