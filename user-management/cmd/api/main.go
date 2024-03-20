package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/pkg/database"
	"user-management/pkg/repository"
	"user-management/pkg/routes"
	"user-management/pkg/usecase"
)

func main() {
	fmt.Println("started user management service!")

	db := database.ConnectDB()

	repo := repository.NewRepositories(db)

	usecase := usecase.NewUseCase(repo)

	handler := routes.NewRoutes(usecase)

	srv := &http.Server{
		Addr: ":8081",
		Handler: handler,
	}

	log.Println("server started at port 8081")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
