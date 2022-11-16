package routes

import (
	"literature/handlers"
	"literature/pkg/middleware"
	"literature/pkg/mysql"
	"literature/repositories"

	"github.com/gorilla/mux"
)

func LiteraturRoutes(r *mux.Router) {
	LiteraturRepository := repositories.RepositoryLiteratur(mysql.DB)

	h := handlers.HandlerLiteratur(LiteraturRepository)

	r.HandleFunc("/literaturs", h.FindLiteraturs).Methods("GET")
	r.HandleFunc("/literaturs/approve", h.FindLiteratursApprove).Methods("GET")

	r.HandleFunc("/literaturs/user/{userId}", middleware.Auth(h.GetLiteraturByUserID)).Methods("GET")

	r.HandleFunc("/literatur", middleware.Auth(middleware.UploadCover(middleware.UploadPDF(h.CreateLiteratur)))).Methods("POST")

	r.HandleFunc("/literatur/{id}", h.GetLiteratur).Methods("GET")

	r.HandleFunc("/literatur/{id}", middleware.Auth(h.DeleteLiteratur)).Methods("DELETE")
	r.HandleFunc("/literatur/{id}", middleware.Auth(middleware.UploadImage(h.UpdateLiteratur))).Methods("PATCH")

	// r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.UploadPost(middleware.UploadPost2(middleware.UploadPost3(middleware.UploadPost4(middleware.UploadPost5(h.UpdateTransaction))))))).Methods("PATCH")
}
