package response

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"net/http"
)

// ErrorResponse - error information to response
type ErrorResponse struct {
	Message string `json:"message" example:"internal server error"`
}

// HTTPErrorResponse - encapsulate http error information to response
type HTTPErrorResponse struct {
	StatusCode int            `json:"statusCode" example:"500"`
	Error      *ErrorResponse `json:"error,omitempty"`
}

// CreateSuccessResponse - creates a success default response
func CreateSuccessResponse(w http.ResponseWriter, statusCode int, result interface{}) error {
	w.Header().Set("Content-Type", constants.JSONContentType)
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(result)
}

// CreateErrorResponse - creates a error default response
func CreateErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) error {
	w.Header().Set("Content-Type", constants.JSONContentType)
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(HTTPErrorResponse{
		Error: &ErrorResponse{
			Message: errorMsg,
		},
		StatusCode: statusCode,
	})
}
