package routes

import (
	"literature/handlers"
	"literature/pkg/middleware"
	"literature/pkg/mysql"
	"literature/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {

	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	// r.HandleFunc("/user", h.GetUser).Methods("GET")

	// r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", middleware.Auth(h.GetUser)).Methods("GET")

	r.HandleFunc("/user/{id}", middleware.Auth(middleware.UploadImage(h.UpdateUser))).Methods("PATCH")
}
