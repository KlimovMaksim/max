package handler

import (
	"encoding/json"
	"max/internal/models"
	"max/internal/service"
	"net/http"
	"strconv"
)

type CardHandler struct {
	service *service.CardService
}

func NewCardHandler(service *service.CardService) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		http.Error(w, "user_id не найден в контексте", http.StatusUnauthorized)
		return
	}
	var req models.CardCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	card := &models.Card{
		AccountID: req.AccountID,
		Number:    req.Number,
		Type:      req.Type,
		UserID:    userID,
	}
	if err := h.service.CreateCard(r.Context(), card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CardHandler) GetCards(w http.ResponseWriter, r *http.Request) {
	accountIDStr := r.URL.Query().Get("account_id")
	if accountIDStr == "" {
		http.Error(w, "account_id обязателен", http.StatusBadRequest)
		return
	}
	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Некорректный account_id", http.StatusBadRequest)
		return
	}
	cards, err := h.service.ListCardsByAccount(r.Context(), accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *CardHandler) GetCard(w http.ResponseWriter, r *http.Request) {
	cardIDStr := r.URL.Query().Get("id")
	if cardIDStr == "" {
		http.Error(w, "id обязателен", http.StatusBadRequest)
		return
	}
	cardID, err := strconv.ParseInt(cardIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Некорректный id", http.StatusBadRequest)
		return
	}
	card, err := h.service.GetCardByID(r.Context(), cardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
