package student

import (
	"net/http"
	"strconv"

	"github.com/thissidemayur/1-golang-studentsApi/internal/storage"
	"github.com/thissidemayur/1-golang-studentsApi/internal/utils/response"
)

func DeleteStudentById(storage storage.Storage) http.HandlerFunc {

	return func (w http.ResponseWriter, r * http.Request) {
		id:=r.PathValue("id")
		intId,err :=strconv.ParseInt(id,10,64)
		if err != nil {
			response.WriteJsonResponse(w,http.StatusBadRequest,response.GeneralError(err))
			return
		}
		err = storage.DeleteStudentById(intId)
		if err != nil {
			response.WriteJsonResponse(w,http.StatusInternalServerError,response.GeneralError(err))
			return
		}
		response.WriteJsonResponse(w,http.StatusOK,map[string]string{"message":"Student deleted successfully"})
	}

}