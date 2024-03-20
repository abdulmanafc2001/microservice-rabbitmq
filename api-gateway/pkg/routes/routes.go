package routes

import (
	helper "api-gateway/pkg/helpers"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user-signup", userSignup)

	mux.HandleFunc("GET /get-allusers", getAllUsers)

	return mux
}

func userSignup(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		First_Name   string `json:"firstname"`
		Last_Name    string `json:"lastname"`
		User_Name    string `json:"username"`
		Password     string `json:"password"`
		Email        string `json:"email"`
		Phone_Number string `json:"phonenumber"`
		Referal_Code string `json:"referalcode"`
	}
	var user payload
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ErrorResponse(w, err)
		return
	}

	jsonData, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "http://localhost:8081/create-user", bytes.NewBuffer(jsonData))

	if err != nil {
		helper.ErrorResponse(w, err)
		return
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		helper.ErrorResponse(w, err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 201 {
		helper.ErrorResponse(w, errors.New("something went wrong please try again"))
		return
	}

	helper.JSONResponse(w, http.StatusCreated, "successfully created user", nil, false)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://localhost:8081/get-allusers")

	if err != nil {
		helper.ErrorResponse(w, err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		helper.ErrorResponse(w, err)
		return
	}

	type JSONResponse struct {
		Error      bool   `json:"error"`
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
		Data       any    `json:"data"`
	}
	var data JSONResponse
	if err := json.Unmarshal(body, &data); err != nil {
		helper.ErrorResponse(w, err)
		return
	}

	helper.JSONResponse(w, res.StatusCode, data.Message, data.Data, false)
}
