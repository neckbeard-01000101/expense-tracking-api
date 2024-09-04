package main

import (
	"expense-tracking/config"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gofor-little/env"
)

func main() {
	err := env.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running and listening to incoming requests ..."))
	})
	config.InitDB()
	fmt.Println("server is running on port :8000")
	http.ListenAndServe(":8000", r)
}
