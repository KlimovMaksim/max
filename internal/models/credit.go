package models

import (
	"time"
	"errors"
)

// Credit представляет модель кредита
type Credit struct {
	ID              int64     `json:"id" db:"id"`
	UserID          int64     `json:"user_id" db:"user_id"`
	AccountID       int64     `json:"account_id" db:"account_id"`
	Amount          float64   `json:"amount" db:"amount"`
	InterestRate    float64   `json:"interest_rate" db:"interest_rate"`
	Term            int       `json:"term" db:"term"` // Срок в месяцах
	MonthlyPayment  float64   `json:"monthly_payment" db:"monthly_payment"`
	RemainingAmount float64   `json:"remaining_amount" db:"remaining_amount"`
	Status          string    `json:"status" db:"status"` // ACTIVE, CLOSED, OVERDUE
	StartDate       time.Time `json:"start_date" db:"start_date"`
	EndDate         time.Time `json:"end_date" db:"end_date"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// CreditCreateRequest представляет запрос на создание кредита
type CreditCreateRequest struct {
	AccountID    int64   `json:"account_id"`
	Amount       float64 `json:"amount"`
	Term         int     `json:"term"` // Срок в месяцах
}

// CreditResponse представляет ответ с данными кредита
type CreditResponse struct {
	ID              int64     `json:"id"`
	AccountID       int64     `json:"account_id"`
	Amount          float64   `json:"amount"`
	InterestRate    float64   `json:"interest_rate"`
	Term            int       `json:"term"`
	MonthlyPayment  float64   `json:"monthly_payment"`
	RemainingAmount float64   `json:"remaining_amount"`
	Status          string    `json:"status"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	CreatedAt       time.Time `json:"created_at"`
}

// CreditListResponse представляет ответ со списком кредитов
type CreditListResponse struct {
	Credits []CreditResponse `json:"credits"`
	Total   int              `json:"total"`
}

// CreditLoadResponse представляет ответ с кредитной нагрузкой
type CreditLoadResponse struct {
	TotalDebt           float64 `json:"total_debt"`
	MonthlyPayments     float64 `json:"monthly_payments"`
	DebtToIncomeRatio   float64 `json:"debt_to_income_ratio"`
	PaymentToIncomeRatio float64 `json:"payment_to_income_ratio"`
	RiskLevel           string  `json:"risk_level"` // LOW, MEDIUM, HIGH
}

// Validate проверяет корректность данных запроса на создание кредита
func (c *CreditCreateRequest) Validate() error {
	if c.Amount <= 0 {
		return errors.New("сумма кредита должна быть положительной")
	}
	
	if c.Term < 1 || c.Term > 360 {
		return errors.New("срок кредита должен быть от 1 до 360 месяцев")
	}
	
	return nil
}