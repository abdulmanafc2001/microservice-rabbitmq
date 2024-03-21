package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-management/pkg/helper"
	"user-management/pkg/models"
	"user-management/pkg/usecase"
)

type JSONResponse struct {
	Error      bool   `json:"error"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type Routes struct {
	Usecase usecase.Usecase
}

func NewRoutes(useCase usecase.Usecase) http.Handler {
	routes := Routes{
		Usecase: useCase,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /create-user", routes.CreateUser)
	mux.HandleFunc("GET /get-allusers", routes.GetUsers)
	mux.HandleFunc("GET /get-user/{id}", routes.GetUserById)
	mux.HandleFunc("PUT /update-user/{id}", routes.UpdateUserById)
	mux.HandleFunc("DELETE /delete-user/{id}", routes.DeleteUserById)
	return mux
}

func (ro *Routes) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := ro.Usecase.CreateUser(user)
	if err != nil {
		resp := JSONResponse{
			Error:      true,
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("failed to create user: %s", err.Error()),
			Data:       nil,
		}
		helper.JSONResponse(w, resp.StatusCode, resp)
		return
	}
	resp := JSONResponse{
		Error:      false,
		StatusCode: 201,
		Message:    "successfully created user",
		Data:       user,
	}
	helper.JSONResponse(w, resp.StatusCode, resp)
	fmt.Println("user created")
}

func (ro *Routes) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ro.Usecase.GetUsers()
	if err != nil {
		resp := JSONResponse{
			Error:      true,
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("failed to get users: %s", err.Error()),
			Data:       nil,
		}
		helper.JSONResponse(w, resp.StatusCode, resp)
		return
	}

	resp := JSONResponse{
		Error:      false,
		StatusCode: http.StatusOK,
		Message:    "user list",
		Data:       users,
	}
	helper.JSONResponse(w, resp.StatusCode, resp)
}

func (ro *Routes) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
	idInt, _ := strconv.Atoi(id)

	user, err := ro.Usecase.GetUserById(idInt)
	if err != nil {
		resp := JSONResponse{
			Error:      true,
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get data",
			Data:       nil,
		}
		helper.JSONResponse(w, http.StatusBadRequest, resp)
		return
	}

	resp := JSONResponse{
		Error:      false,
		StatusCode: http.StatusOK,
		Message:    "user",
		Data:       user,
	}
	helper.JSONResponse(w, http.StatusOK, resp)
}

func (ro *Routes) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	usrId, _ := strconv.Atoi(id)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp := JSONResponse{
			Error:      true,
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get body",
			Data:       nil,
		}
		helper.JSONResponse(w, resp.StatusCode, resp)
		return
	}

	err = ro.Usecase.UpdateUserById(usrId, user)
	if err != nil {
		resp := JSONResponse{
			Error:      true,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		helper.JSONResponse(w, resp.StatusCode, resp)
		return
	}

	resp := JSONResponse{
		Error:      false,
		StatusCode: http.StatusOK,
		Message:    "Successfully edited user data",
		Data:       nil,
	}
	helper.JSONResponse(w, resp.StatusCode, resp)
}

func (ro *Routes) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	usrId, _ := strconv.Atoi(id)

	err := ro.Usecase.DeleteUserById(usrId)
	if err != nil {
		resp := JSONResponse{
			Error:      true,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		helper.JSONResponse(w, resp.StatusCode, resp)
		return
	}

	resp := JSONResponse{
		Error:      false,
		StatusCode: 200,
		Message:    "successfully deleted user",
		Data:       nil,
	}
	helper.JSONResponse(w, resp.StatusCode, resp)
}
