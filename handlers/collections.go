package handlers

import (
	"encoding/json"
	collectiondto "literature/dto/collections"
	dto "literature/dto/result"
	"literature/models"
	"literature/repositories"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// var path_file_collection = "http://localhost:5000/uploads/"

type handlerCollection struct {
	CollectionRepository repositories.CollectionRepository
}

func HandlerCollection(CollectionRepository repositories.CollectionRepository) *handlerCollection {
	return &handlerCollection{CollectionRepository}

}

func (h *handlerCollection) CreateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["user_id"].(float64))

	// var request bookmarkdto.BookmarkRequest
	var request collectiondto.CollectionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	collection := models.Collection{
		UserID:      userId,
		LiteraturID: request.LiteraturID,
	}

	data, err := h.CollectionRepository.CreateCollection(collection)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	collection, _ = h.CollectionRepository.GetCollection(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCollection(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseCollection(u models.Collection) models.Collection {
	return models.Collection{
		ID:          u.ID,
		UserID:      u.UserID,
		User:        u.User,
		LiteraturID: u.LiteraturID,
		Literatur:   u.Literatur,
	}
}

func (h *handlerCollection) FindCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collections, err := h.CollectionRepository.FindCollection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// for i, p := range bookmarks {
	// 	bookmarks[i].Image = os.Getenv("PATH_FILE") + p.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: collections}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCollection) GetCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var collection models.Collection

	collection, err := h.CollectionRepository.GetCollection(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: collection}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCollection) GetCollectionByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["user_id"].(float64))

	var collections []models.Collection
	// literaturs, err := h.LiteraturRepository.GetLiteraturByUserID(id)
	collections, err := h.CollectionRepository.GetCollectionByUserID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: collections}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCollection) DeleteCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	collection, err := h.CollectionRepository.GetCollection(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CollectionRepository.DeleteCollection(collection)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCollection(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCollection) GetCollectionByLiteratur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var collection []models.Collection
	collection, err := h.CollectionRepository.GetCollectionByLiteratur(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: collection}
	json.NewEncoder(w).Encode(response)
}
