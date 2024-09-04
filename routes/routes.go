package routes

import (
	"expense-tracking/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running and listening to incoming requests ..."))
	})
	r.Post("/register", controllers.Register)
	r.Post("/login", controllers.Login)
}
