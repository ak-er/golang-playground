package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, message interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(message)
}

func GeneralMessage(status string, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMessage []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMessage = append(errMessage, fmt.Sprintf("field %s is required field", err.Field()))
		default:
			errMessage = append(errMessage, err.Error())
		}
	}
	return Response{
		Status:  StatusError,
		Message: strings.Join(errMessage, ", "),
	}
}
