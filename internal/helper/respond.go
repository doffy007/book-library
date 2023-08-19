package helper

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{}, status int) {
	respData := data

	switch v := data.(type) {
	case ErrArgument:
		status = http.StatusBadRequest
		respData = ErrorResponse{ErrorMessage: v.Unwrap().Error()}
	case error:
		if http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		} else {
			respData = ErrorResponse{ErrorMessage: v.Error()}
		}
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if respData != nil {
		if err := json.NewEncoder(w).Encode(respData); err != nil {
			http.Error(w, "Could not encode in JSON", http.StatusBadRequest)
			return
		}
	}
}
