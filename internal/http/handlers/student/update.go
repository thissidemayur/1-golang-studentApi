package student

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/thissidemayur/1-golang-studentsApi/internal/storage"
	"github.com/thissidemayur/1-golang-studentsApi/internal/types"
	"github.com/thissidemayur/1-golang-studentsApi/internal/utils/response"
)

func UpdateStudentById(storage storage.Storage) http.HandlerFunc {
	return func (w http.ResponseWriter, r * http.Request) {
		idStr := r.PathValue("id")
		intId,err:=strconv.ParseInt(idStr,10,64)
		if err != nil {
			response.WriteJsonResponse(w,http.StatusBadRequest,response.GeneralError(err))
			return
		}

		// decode request
		var dto types.UpdateStudent
		err = json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			response.WriteJsonResponse(w,http.StatusBadRequest,response.GeneralError(err))
			return
		}
		// validation
		if dto.Email == nil && dto.Name == nil && dto.RollNo == nil {
			
			response.WriteJsonResponse(w,http.StatusBadRequest,response.GeneralError(errors.New("At least one field must be provided to update")))
			return

		}

		// update student
		updatedStudent,err := storage.UpdateStudentById(intId,dto)
		if err != nil {
			response.WriteJsonResponse(w,http.StatusInternalServerError,response.GeneralError(err))
			return
		}

		response.WriteJsonResponse(w,http.StatusOK,updatedStudent)
	}

}