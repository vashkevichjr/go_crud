package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vashkevichjr/go_crud/internal/repository"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

type Request struct {
	Number int `json:"number"`
}

type Response struct {
	SortedNums []int `json:"sortedNums"`
}

func (h *Handler) SaveAndGetNumber(w http.ResponseWriter, r *http.Request) {
	req := Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.SaveNumber(r.Context(), req.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{}
	response.SortedNums, err = h.repo.GetSortedNums(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}

}
