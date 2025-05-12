package models

import (
	"errors"
	"regexp"
	"time"
	"max/pkg/req"
)

// User представляет модель пользователя
type User struct {
	ID           int64     `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"password_hash" db:"password_hash"` // Добавление PasswordHash в структуру User
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// UserRegisterRequest представляет запрос на регистрацию
type UserRegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

// UserLoginRequest представляет запрос на аутентификацию
type UserLoginRequest struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}

// UserResponse представляет ответ с данными пользователя
type UserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

// AuthResponse представляет ответ с токеном аутентификации
type AuthResponse struct {
	Token string       `json:"token"`
}

// Validate проверяет корректность данных пользователя
func (u *UserRegisterRequest) Validate() error {
	if err := req.IsValid(u); err != nil {
		return errors.New("некорректные данные")
	}
	
	// Проверка имени пользователя
	if len(u.Username) < 3 || len(u.Username) > 50 {
		return errors.New("имя пользователя должно содержать от 3 до 50 символов")
	}

	// Проверка email
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("некорректный формат email")
	}

	// Проверка пароля
	if len(u.Password) < 8 {
		return errors.New("пароль должен содержать не менее 8 символов")
	}

	return nil
}

// ToUser преобразует запрос регистрации в модель пользователя
func (u *UserRegisterRequest) ToUser() *User {
	return &User{
		Username:  u.Username,
		Email:     u.Email,
		PasswordHash:  u.Password, // Пароль будет хешироваться в сервисе
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// ToResponse преобразует модель пользователя в ответ API
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
	}
}
