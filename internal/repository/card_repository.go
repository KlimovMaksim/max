package repository

import (
	"max/internal/models"
	"context"
	"max/pkg/database"
)

type CardRepository struct {
	db *database.Db
}

func NewCardRepository(db *database.Db) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) Create(ctx context.Context, card *models.Card) error {
	query := "INSERT INTO cards (account_id, number, expiry, cvv, created_at) VALUES (?, ?, ?, ?, ?)"
	result := r.db.Exec(query, card.AccountID, card.Number, card.Expiry, card.CVV, card.CreatedAt)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CardRepository) GetByID(ctx context.Context, id int64) (*models.Card, error) {
	query := "SELECT id, account_id, number, expiry, cvv, created_at FROM cards WHERE id = ?"
	row := r.db.Raw(query, id).Row()
	var c models.Card
	err := row.Scan(&c.ID, &c.AccountID, &c.Number, &c.Expiry, &c.CVV, &c.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CardRepository) ListByAccount(ctx context.Context, accountID int64) ([]models.Card, error) {
	query := "SELECT id, account_id, number, expiry, cvv, created_at FROM cards WHERE account_id = ?"
	rows, err := r.db.Raw(query, accountID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []models.Card
	for rows.Next() {
		var c models.Card
		if err := rows.Scan(&c.ID, &c.AccountID, &c.Number, &c.Expiry, &c.CVV, &c.CreatedAt); err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, nil
}
