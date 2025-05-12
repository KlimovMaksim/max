package models

import (
	"time"
)

// TransactionType представляет тип транзакции
type TransactionType string

const (
	TransactionTypeDeposit  TransactionType = "DEPOSIT"
	TransactionTypeWithdraw TransactionType = "WITHDRAW"
	TransactionTypeTransfer TransactionType = "TRANSFER"
	TransactionTypePayment  TransactionType = "PAYMENT"
	TransactionTypeFee      TransactionType = "FEE"
)

// Transaction представляет модель транзакции
type Transaction struct {
	ID            int64           `json:"id" db:"id"`
	UserID        int64           `json:"user_id" db:"user_id"`
	FromAccountID *int64          `json:"from_account_id,omitempty" db:"from_account_id"`
	ToAccountID   *int64          `json:"to_account_id,omitempty" db:"to_account_id"`
	Type          TransactionType `json:"type" db:"type"`
	Amount        float64         `json:"amount" db:"amount"`
	Description   string          `json:"description" db:"description"`
	Status        string          `json:"status" db:"status"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	AccountID     *int64          `json:"account_id,omitempty" db:"account_id"` // Добавление AccountID в структуру Transaction
}

// TransactionResponse представляет ответ с данными транзакции
type TransactionResponse struct {
	ID            int64           `json:"id"`
	FromAccountID *int64          `json:"from_account_id,omitempty"`
	ToAccountID   *int64          `json:"to_account_id,omitempty"`
	Type          TransactionType `json:"type"`
	Amount        float64         `json:"amount"`
	Description   string          `json:"description"`
	Status        string          `json:"status"`
	CreatedAt     time.Time       `json:"created_at"`
}

// TransactionListResponse представляет ответ со списком транзакций
type TransactionListResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
	Total        int                   `json:"total"`
}

// AnalyticsResponse представляет ответ с аналитикой транзакций
type AnalyticsResponse struct {
	TotalIncome       float64                 `json:"total_income"`
	TotalExpense      float64                 `json:"total_expense"`
	NetChange         float64                 `json:"net_change"`
	Categories        map[string]CategoryStat `json:"categories"`
	MonthlyStats      []MonthlyStat           `json:"monthly_stats"`
	TotalBalance      float64                 `json:"total_balance"`      // Добавление TotalBalance в структуру AnalyticsResponse
	TotalCredits      float64                 `json:"total_credits"`      // Добавление TotalCredits в структуру AnalyticsResponse
	TotalTransactions int                     `json:"total_transactions"` // Добавление TotalTransactions в структуру AnalyticsResponse
}

// CategoryStat представляет статистику по категории
type CategoryStat struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}

// MonthlyStat представляет статистику за месяц
type MonthlyStat struct {
	Month   string  `json:"month"` // Формат "YYYY-MM"
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}
