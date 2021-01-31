package response

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/util/constants"
	"net/http"
)

func CreateSuccessResponse(w http.ResponseWriter, statusCode int, result interface{}) error {
	w.Header().Set("Content-Type", constants.JsonContentType)
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(result)
}

func CreateErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) error {
	w.Header().Set("Content-Type", constants.JsonContentType)
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(errorMsg)
}
