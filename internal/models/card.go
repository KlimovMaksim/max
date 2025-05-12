package models

import (
	"errors"
	"time"
)

// Card представляет модель банковской карты
type Card struct {
	ID             int64     `json:"id" db:"id"`
	AccountID      int64     `json:"account_id" db:"account_id"`
	UserID         int64     `json:"user_id" db:"user_id"`
	Number         string    `json:"-" db:"number_encrypted"`      // Зашифрованный номер карты
	NumberHMAC     []byte    `json:"-" db:"number_hmac"`           // HMAC для проверки целостности
	ExpiryDate     []byte    `json:"-" db:"expiry_date_encrypted"` // Зашифрованная дата истечения
	ExpiryDateHMAC []byte    `json:"-" db:"expiry_date_hmac"`      // HMAC для проверки целостности
	CVV            []byte    `json:"-" db:"cvv_hash"`              // Хеш CVV
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Expiry         time.Time `json:"expiry"`         // Добавление Expiry в структуру Card
	Type           string    `json:"type" db:"type"` // Тип карты (добавлено)
}

// CardCreateRequest представляет запрос на создание карты
type CardCreateRequest struct {
	AccountID int64  `json:"account_id"`
	Number    string `json:"number"`
	Type      string `json:"type"`
}

// Validate проверяет корректность данных запроса на создание карты
func (c *CardCreateRequest) Validate() error {
	if c.AccountID <= 0 {
		return errors.New("account_id должен быть положительным")
	}
	if len(c.Number) < 12 || len(c.Number) > 19 {
		return errors.New("номер карты должен содержать от 12 до 19 цифр")
	}
	if c.Type == "" {
		return errors.New("тип карты не должен быть пустым")
	}
	return nil
}

// CardResponse представляет ответ с данными карты
type CardResponse struct {
	ID         int64     `json:"id"`
	AccountID  int64     `json:"account_id"`
	Number     string    `json:"number"`      // Маскированный номер (только последние 4 цифры)
	ExpiryDate string    `json:"expiry_date"` // Формат MM/YY
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
}

// CardDetailResponse представляет ответ с полными данными карты (для владельца)
type CardDetailResponse struct {
	ID         int64     `json:"id"`
	AccountID  int64     `json:"account_id"`
	Number     string    `json:"number"`      // Полный номер (только для владельца)
	ExpiryDate string    `json:"expiry_date"` // Формат MM/YY
	CVV        string    `json:"cvv"`         // CVV (только для владельца)
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
}
