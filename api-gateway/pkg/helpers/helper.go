package helper

import (
	js "encoding/json"
	"log"
	"net/http"
)

type JSON struct {
	Err        bool   `json:"error"`
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, err error, stCode ...int) {
	statuscode := http.StatusBadRequest
	if len(stCode) > 0 {
		statuscode = stCode[0]
	}
	JSONResponse(w, statuscode, err.Error(), nil, true)
}

func JSONResponse(w http.ResponseWriter, statuscode int, message string, data any, errExist bool) {
	json := JSON{
		Err:        errExist,
		StatusCode: statuscode,
		Message:    message,
		Data:       data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	err := js.NewEncoder(w).Encode(&json)
	if err != nil {
		log.Fatal(err)
	}
}
