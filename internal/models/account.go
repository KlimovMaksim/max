package models

import (
	"time"
	"errors"
)

// AccountType представляет тип банковского счета
type AccountType string

const (
	AccountTypeDebit  AccountType = "DEBIT"
	AccountTypeCredit AccountType = "CREDIT"
)

// Account представляет модель банковского счета
type Account struct {
	ID        int64       `json:"id" db:"id"`
	UserID    int64       `json:"user_id" db:"user_id"`
	Number    string      `json:"number" db:"number"`
	Type      AccountType `json:"type" db:"type"`
	Balance   float64     `json:"balance" db:"balance"`
	Currency  string      `json:"currency" db:"currency"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}

// AccountCreateRequest представляет запрос на создание счета
type AccountCreateRequest struct {
	Type     AccountType `json:"type"`
	Currency string      `json:"currency"`
}

// DepositRequest представляет запрос на пополнение счета
type DepositRequest struct {
	Amount float64 `json:"amount"`
}

// WithdrawRequest представляет запрос на снятие средств
type WithdrawRequest struct {
	Amount float64 `json:"amount"`
}

// TransferRequest представляет запрос на перевод средств
type TransferRequest struct {
	FromAccountID int64   `json:"from_account_id"`
	ToAccountID   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
}

// PredictBalanceRequest представляет запрос на прогноз баланса
type PredictBalanceRequest struct {
	Days int `json:"days"`
}

// PredictBalanceResponse представляет ответ с прогнозом баланса
type PredictBalanceResponse struct {
	CurrentBalance float64                `json:"current_balance"`
	PredictedBalance float64              `json:"predicted_balance"`
	Transactions []PredictedTransaction   `json:"transactions"`
}

// PredictedTransaction представляет прогнозируемую транзакцию
type PredictedTransaction struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"` // "income" или "expense"
}

// Validate проверяет корректность данных запроса на создание счета
func (a *AccountCreateRequest) Validate() error {
	if a.Type != AccountTypeDebit && a.Type != AccountTypeCredit {
		return errors.New("неверный тип счета")
	}

	if a.Currency != "RUB" {
		return errors.New("поддерживается только валюта RUB")
	}

	return nil
}

// Validate проверяет корректность данных запроса на пополнение счета
func (d *DepositRequest) Validate() error {
	if d.Amount <= 0 {
		return errors.New("сумма пополнения должна быть положительной")
	}
	return nil
}

// Validate проверяет корректность данных запроса на снятие средств
func (w *WithdrawRequest) Validate() error {
	if w.Amount <= 0 {
		return errors.New("сумма снятия должна быть положительной")
	}
	return nil
}

// Validate проверяет корректность данных запроса на перевод средств
func (t *TransferRequest) Validate() error {
	if t.Amount <= 0 {
		return errors.New("сумма перевода должна быть положительной")
	}
	
	if t.FromAccountID == t.ToAccountID {
		return errors.New("счета отправителя и получателя должны различаться")
	}
	
	return nil
}

// Validate проверяет корректность данных запроса на прогноз баланса
func (p *PredictBalanceRequest) Validate() error {
	if p.Days <= 0 {
		return errors.New("количество дней должно быть положительным")
	}
	
	if p.Days > 365 {
		return errors.New("максимальный период прогноза - 365 дней")
	}
	
	return nil
}