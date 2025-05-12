package repository

import (
	"context"
	"max/pkg/database"
	"max/internal/models"
)

type TransactionRepository struct {
	db *database.Db
}

func NewTransactionRepository(db *database.Db) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *models.Transaction) error {
	query := "INSERT INTO transactions (account_id, amount, type, description, created_at) VALUES (?, ?, ?, ?, ?)"
	result := r.db.Exec(query, transaction.AccountID, transaction.Amount, transaction.Type, transaction.Description, transaction.CreatedAt)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TransactionRepository) GetByID(ctx context.Context, id int64) (*models.Transaction, error) {
	query := "SELECT id, account_id, amount, type, description, created_at FROM transactions WHERE id = ?"
	row := r.db.Raw(query, id).Row()
	var t models.Transaction
	err := row.Scan(&t.ID, &t.AccountID, &t.Amount, &t.Type, &t.Description, &t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TransactionRepository) ListByAccount(ctx context.Context, accountID int64) ([]models.Transaction, error) {
	query := "SELECT id, account_id, amount, type, description, created_at FROM transactions WHERE account_id = ?"
	rows, err := r.db.Raw(query, accountID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.AccountID, &t.Amount, &t.Type, &t.Description, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
