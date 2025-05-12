package models

import (
	"time"
)

// PaymentSchedule представляет график платежей по кредиту
// Один платеж — одна запись
// Все суммы в RUB
// Статус: SCHEDULED, PAID, OVERDUE
// Penalty — начисленная пеня (если есть)
type PaymentSchedule struct {
	ID        int64      `json:"id" db:"id"`
	CreditID  int64      `json:"credit_id" db:"credit_id"`
	AccountID int64      `json:"account_id" db:"account_id"`
	DueDate   time.Time  `json:"due_date" db:"due_date"`
	Amount    float64    `json:"amount" db:"amount"`
	Penalty   float64    `json:"penalty" db:"penalty"`
	Status    string     `json:"status" db:"status"`
	PaidAt    *time.Time `json:"paid_at,omitempty" db:"paid_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

// PaymentScheduleResponse — ответ с данными по платежу
// Используется для отображения графика клиенту
type PaymentScheduleResponse struct {
	ID      int64      `json:"id"`
	DueDate time.Time  `json:"due_date"`
	Amount  float64    `json:"amount"`
	Penalty float64    `json:"penalty"`
	Status  string     `json:"status"`
	PaidAt  *time.Time `json:"paid_at,omitempty"`
}

// PaymentScheduleListResponse — список платежей по кредиту
type PaymentScheduleListResponse struct {
	Payments []PaymentScheduleResponse `json:"payments"`
	Total    int                       `json:"total"`
}
