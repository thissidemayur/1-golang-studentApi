package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	StatusOk = "ok"
	StatusError = "error"
)

func WriteJsonResponse(w http.ResponseWriter,status int,data interface{}) error {

	// set header
	w.Header().Set("Content-Type","application/json")

	// set status code
	w.WriteHeader(status)

	// send response
	return json.NewEncoder(w).Encode(data)

}


// Error struct to manage error
type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func GeneralError(err error) Response {
	return Response{
		Status:  StatusError,
		Message: err.Error(),
	}
}

// validation error by validator package
func ValidationError(errs validator.ValidationErrors) Response  {
	fmt.Printf("Validation errors: %+v\n", errs)

	var errMsgs []string;
	for _,err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("The field %s is required.", err.Field()))

		case "email":
			errMsgs = append(errMsgs, fmt.Sprintf("The field %s must be a valid email address.", err.Field()))

		case "min":
			errMsgs = append(errMsgs, fmt.Sprintf("The field %s must be at least %s characters long.", err.Field(), err.Param()))

		case "max":
			errMsgs = append(errMsgs, fmt.Sprintf("The field %s must be at most %s characters long.", err.Field(), err.Param()))
			
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("The field %s is invalid.", err.Field()))
		}
		
	}

	return Response{
		Status:  StatusError,
		Message: strings.Join(errMsgs,", "),
	}
}