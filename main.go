package main

import (
	"expense-tracking/config"
	"expense-tracking/routes"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gofor-little/env"
	"log"
	"net/http"
)

func main() {
	err := env.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	routes.RegisterRoutes(r)
	fmt.Println("server is running on port :8000")
	http.ListenAndServe(":8000", r)
}
