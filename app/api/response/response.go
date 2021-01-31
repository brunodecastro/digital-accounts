package response

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type ApiHttpResponseSuccess struct {
	Error      *ErrorResponse `json:"error,omitempty"`
	StatusCode int            `json:"statusCode"`
	Result     interface{}    `json:"result,omitempty"`
}

func CreateSuccessResponse(w http.ResponseWriter, statusCode int, result interface{}) error {
	w.Header().Set("Content-Type", constants.JsonContentType)
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(ApiHttpResponseSuccess{
		StatusCode: statusCode,
		Result:     result,
	})
}

func CreateErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) error {
	w.Header().Set("Content-Type", constants.JsonContentType)
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(ApiHttpResponseSuccess{
		Error: &ErrorResponse{
			Message: errorMsg,
		},
		StatusCode: statusCode,
	})
}
