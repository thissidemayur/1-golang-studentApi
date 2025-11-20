package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/thissidemayur/1-golang-studentsApi/internal/types"
	"github.com/thissidemayur/1-golang-studentsApi/internal/utils/response"
)


func New(w http.ResponseWriter, r *http.Request) {
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

	w.Write([]byte("Welcome to Student Handler API"))
	response.WriteJsonResponse(w,http.StatusCreated,map[string]string{
		"success":"true",
		"message":"Student created successfully",
	})

}


/*
func New(){}  || func Create() {}  -> convention for handler function name
*/
