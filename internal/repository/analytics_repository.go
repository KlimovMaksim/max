package repository

import (
	"context"
	"max/pkg/database"
	"max/internal/models"
)

type AnalyticsRepository struct {
	db *database.Db
}

func NewAnalyticsRepository(db *database.Db) *AnalyticsRepository {
	return &AnalyticsRepository{db: db}
}

func (r *AnalyticsRepository) GetAnalytics(ctx context.Context, userID int64) (*models.AnalyticsResponse, error) {
	query := "SELECT total_balance, total_credits, total_transactions FROM analytics WHERE user_id = ?"
	row := r.db.Raw(query, userID).Row()
	var resp models.AnalyticsResponse
	err := row.Scan(&resp.TotalBalance, &resp.TotalCredits, &resp.TotalTransactions)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *AnalyticsRepository) GetCreditLoad(ctx context.Context, userID int64) (float64, error) {
	query := "SELECT SUM(amount) FROM credits WHERE user_id = ? AND status = 'active'"
	row := r.db.Raw(query, userID).Row()
	var load float64
	err := row.Scan(&load)
	if err != nil {
		return 0, err
	}
	return load, nil
}
