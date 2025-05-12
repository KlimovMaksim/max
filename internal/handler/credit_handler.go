package handler

import (
	"encoding/json"
	"max/internal/models"
	"max/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CreditHandler struct {
	service *service.CreditService
}

func NewCreditHandler(service *service.CreditService) *CreditHandler {
	return &CreditHandler{service: service}
}

// Создание кредита
func (h *CreditHandler) CreateCredit(w http.ResponseWriter, r *http.Request) {
	var req models.CreditCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный формат запроса", http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.WriteHeader(http.StatusCreated)

}

// Получение всех кредитов пользователя
func (h *CreditHandler) GetCredits(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("user_id").(int64)
	if !ok {
		http.Error(w, "user_id не найден в контексте", http.StatusUnauthorized)
		return
	}
}

// Получение информации о конкретном кредите
func (h *CreditHandler) GetCredit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Некорректный id", http.StatusBadRequest)
		return
	}
	credit, err := h.service.GetCreditByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(credit)
}

// Получение графика платежей по кредиту
func (h *CreditHandler) GetPaymentSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Некорректный id", http.StatusBadRequest)
		return
	}
	
}
