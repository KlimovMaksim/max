package repository

import (
	"context"
	"max/internal/models"
	"max/pkg/database"
)

type AccountRepository struct {
	db *database.Db
}

func NewAccountRepository(db *database.Db) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(ctx context.Context, account *models.Account) error {
	query := "INSERT INTO accounts (user_id, balance, created_at) VALUES (?, ?, ?)"
	result := r.db.Exec(query, account.UserID, account.Balance, account.CreatedAt)
	if result.Error != nil {
		return result.Error
		
	}
	return nil
}

func (r *AccountRepository) GetByID(ctx context.Context, id int64) (*models.Account, error) {
	query := "SELECT id, user_id, balance, created_at FROM accounts WHERE id = ?"
	row := r.db.Raw(query, id).Row()
	var a models.Account
	err := row.Scan(&a.ID, &a.UserID, &a.Balance, &a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AccountRepository) ListByUser(ctx context.Context, userID int64) ([]models.Account, error) {
	query := "SELECT id, user_id, balance, created_at FROM accounts WHERE user_id = ?"
	rows, err := r.db.Raw(query, userID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []models.Account
	for rows.Next() {
		var a models.Account
		if err := rows.Scan(&a.ID, &a.UserID, &a.Balance, &a.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}
