package response

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message" example:"internal server error"`
}

type HttpErrorResponse struct {
	StatusCode int            `json:"statusCode" example:"500"`
	Error      *ErrorResponse `json:"error,omitempty"`
}

func CreateSuccessResponse(w http.ResponseWriter, statusCode int, result interface{}) error {
	w.Header().Set("Content-Type", constants.JsonContentType)
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(result)
}

func CreateErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) error {
	w.Header().Set("Content-Type", constants.JsonContentType)
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(HttpErrorResponse{
		Error: &ErrorResponse{
			Message: errorMsg,
		},
		StatusCode: statusCode,
	})
}
