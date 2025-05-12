package handler

import (
	"encoding/json"
	"max/internal/models"
	"max/internal/service"
	"net/http"

)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Регистрация пользователя
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный формат запроса", http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := req.ToUser()
	if err := h.service.Register(r.Context(), user, req.Password); err != nil {
		if err.Error() == "user already exists" {
			http.Error(w, "Пользователь уже существует", http.StatusConflict)
			return
		}
		http.Error(w, "Ошибка регистрации пользователя", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user.ToResponse())
}

// Логин пользователя
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный формат запроса", http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Логин и пароль обязательны", http.StatusBadRequest)
		return
	}
	token, err := h.service.Authenticate(r.Context(), req.Username, req.Password)
	if err != nil {
		if err.Error() == "неверный пароль" {
			http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
			return
		}
	
		http.Error(w, "Ошибка аутентификации", http.StatusInternalServerError)
		return
	}
	
	resp := models.AuthResponse{
		Token: token,
		
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
