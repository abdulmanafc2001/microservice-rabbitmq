package helper

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data any) error {
	byteData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(byteData)
	return nil
}

func HashPassword(password string) (string, error) {
	passByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", errors.New("failed to hash password: " + err.Error())
	}
	return string(passByte), nil
}
