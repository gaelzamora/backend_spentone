package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gaelzamora/spent-one/internal/application"
	"github.com/gaelzamora/spent-one/internal/domain"
	"github.com/gorilla/mux"
)

type SpentHandler struct {
	service *application.SpentService
}

func NewSpentHandler(service *application.SpentService) *SpentHandler {
	return &SpentHandler{service: service}
}

func (h *SpentHandler) CreateSpent(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var spent domain.Spent
	spent.UserID = userID
	err := json.NewDecoder(r.Body).Decode(&spent)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	spent, err = h.service.CreateSpent(&spent)

	if err != nil {
		http.Error(w, "Error creando gasto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SpentHandler) GetSpents(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	spents, err := h.service.GetSpents(userID)

	if err != nil {
		http.Error(w, "Error obteniendo gastos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spents)
}

func (h *SpentHandler) GetSpent(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	vars := mux.Vars(r)
	spentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	spent, err := h.service.GetSpent(userID, uint(spentID))

	if err != nil {
		http.Error(w, "Error obteniendo gasto", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spent)
}

func (h *SpentHandler) DeleteSpent(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	vars := mux.Vars(r)
	spentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteSpent(userID, uint(spentID))

	if err != nil {
		http.Error(w, "Error al eliminar el gasto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SpentHandler) UpdateSpent(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	existingSpent, err := h.service.GetSpent(userID, uint(id))

	if err != nil {
		http.Error(w, "Error obteniendo gasto", http.StatusNotFound)
		return
	}

	var updatedData domain.Spent
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	existingSpent.Amount = updatedData.Amount
	existingSpent.Reason = updatedData.Reason

	err = h.service.UpdateSpent(existingSpent)
	if err != nil {
		http.Error(w, "Error actualizando el gasto", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingSpent)
}
