package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	literaturdto "literature/dto/literatur"
	dto "literature/dto/result"
	"literature/models"
	"literature/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerLiteratur struct {
	LiteraturRepository repositories.LiteraturRepository
}

func HandlerLiteratur(LiteraturRepository repositories.LiteraturRepository) *handlerLiteratur {
	return &handlerLiteratur{LiteraturRepository}
}

// -----------------------------------

func (h *handlerLiteratur) CreateLiteratur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	fmt.Println("DATA USER", userInfo)
	user_id := int(userInfo["user_id"].(float64))

	fmt.Println("DATA", user_id)

	coverContext := r.Context().Value("dataFile")
	filepath := coverContext.(string)

	pdfContext := r.Context().Value("dataPDF")
	filename := pdfContext.(string)

	pages, _ := strconv.Atoi(r.FormValue("pages"))

	request := literaturdto.CreateLiteratureRequest{
		UserID:             user_id,
		Title:              r.FormValue("title"),
		PublicationDate:    r.FormValue("publication_date"),
		Pages:              pages,
		ISBN:               r.FormValue("isbn"),
		Author:             r.FormValue("author"),
		Attache:            filename,
		Cover:              filepath,
		Statusverification: "pending",
	}
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "goLiteratur"})

	if err != nil {
		fmt.Println(err.Error())
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// pages, _ := strconv.Atoi(r.FormValue("pages"))

	literatur := models.Literatur{
		Title:              request.Title,
		PublicationDate:    request.PublicationDate,
		Pages:              pages,
		ISBN:               request.ISBN,
		Author:             request.Author,
		Attache:            filename,
		Cover:              resp.SecureURL,
		UserID:             user_id,
		Statusverification: request.Statusverification,
	}

	literatur, err = h.LiteraturRepository.CreateLiteratur(literatur)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	literatur, err = h.LiteraturRepository.GetLiteratur(literatur.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	literatur, _ = h.LiteraturRepository.GetLiteratur(literatur.ID)
	literatur.Attache = os.Getenv("PATH_FILE") + literatur.Attache

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: literatur}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerLiteratur) GetLiteratur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	literatur, err := h.LiteraturRepository.GetLiteratur(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	literatur.Attache = os.Getenv("PATH_FILE") + literatur.Attache

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: literatur}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerLiteratur) DeleteLiteratur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	literatur, err := h.LiteraturRepository.GetLiteratur(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.LiteraturRepository.DeleteLiteratur(literatur, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerLiteratur) FindLiteraturs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	literaturs, err := h.LiteraturRepository.FindLiteraturs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range literaturs {
		literaturs[i].Attache = os.Getenv("PATH_FILE") + p.Attache
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: literaturs}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerLiteratur) FindLiteratursApprove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	literaturs, err := h.LiteraturRepository.FindLiteratursApprove()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range literaturs {
		literaturs[i].Attache = os.Getenv("PATH_FILE") + p.Attache
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: literaturs}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerLiteratur) GetLiteraturByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	var literaturs []models.Literatur
	literaturs, err := h.LiteraturRepository.GetLiteraturByUserID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: literaturs}
	json.NewEncoder(w).Encode(response)
}

func convertResponseLiteratur(u models.Literatur) models.LiteraturResponse {
	return models.LiteraturResponse{
		Title:           u.Title,
		User:            u.User,
		PublicationDate: u.PublicationDate,
		Pages:           u.Pages,
		ISBN:            u.ISBN,
		Author:          u.Author,
		Attache:         u.Attache,
	}
}

func (h *handlerLiteratur) UpdateLiteratur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := literaturdto.UpdateLiteraturRequest{
		Statusverification: r.FormValue("statusverification"),
	}
	literatur := models.Literatur{}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if request.Statusverification != "" {
		literatur.Statusverification = request.Statusverification
	}

	literatur, err := h.LiteraturRepository.UpdateLiteratur(literatur, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: literatur}
	json.NewEncoder(w).Encode(response)
}
