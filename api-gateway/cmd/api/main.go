package main

import (
	"api-gateway/pkg/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("welcome to api gateway")

	handler := routes.Routes()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Fatal(srv.ListenAndServe())
}
