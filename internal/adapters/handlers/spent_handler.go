package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gaelzamora/spent-one/internal/domain"
)

func (h *SpentHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var spent domain.Spent
	err := json.NewDecoder(r.Body).Decode(&spent)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	spent, err := h.service.Create(userID, &spent)

	if err != nil {
		http.Error(w, "Error creando tarea", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(spent)
}