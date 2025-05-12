package repository

import (
	"max/internal/models"
	"context"
	"max/pkg/database"
	"errors"
	"time"
)

type CreditRepository struct {
	db *database.Db
}

func NewCreditRepository(db *database.Db) *CreditRepository {
	return &CreditRepository{db: db}
}

func (r *CreditRepository) Create(ctx context.Context, credit *models.Credit) error {
	query := `INSERT INTO credits (user_id, account_id, amount, interest_rate, term, monthly_payment, remaining_amount, status, start_date, end_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result := r.db.Exec(query,
		credit.UserID,
		credit.AccountID,
		credit.Amount,
		credit.InterestRate,
		credit.Term,
		credit.MonthlyPayment,
		credit.RemainingAmount,
		credit.Status,
		credit.StartDate,
		credit.EndDate,
		time.Now(),
		time.Now(),
	)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CreditRepository) GetByID(ctx context.Context, id int64) (*models.Credit, error) {
	query := `SELECT id, user_id, account_id, amount, interest_rate, term, monthly_payment, remaining_amount, status, start_date, end_date, created_at, updated_at FROM credits WHERE id = ?`
	row := r.db.Raw(query, id).Row()
	var credit models.Credit
	err := row.Scan(
		&credit.ID,
		&credit.UserID,
		&credit.AccountID,
		&credit.Amount,
		&credit.InterestRate,
		&credit.Term,
		&credit.MonthlyPayment,
		&credit.RemainingAmount,
		&credit.Status,
		&credit.StartDate,
		&credit.EndDate,
		&credit.CreatedAt,
		&credit.UpdatedAt,
	)
	if err != nil {
		return nil, errors.New("кредит не найден")
	}
	return &credit, err
}

func (r *CreditRepository) ListByAccount(ctx context.Context, accountID int64) ([]models.Credit, error) {
	query := `SELECT id, user_id, account_id, amount, interest_rate, term, monthly_payment, remaining_amount, status, start_date, end_date, created_at, updated_at FROM credits WHERE account_id = ?`
	rows, err := r.db.Raw(query, accountID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var credits []models.Credit
	for rows.Next() {
		var credit models.Credit
		err := rows.Scan(
			&credit.ID,
			&credit.UserID,
			&credit.AccountID,
			&credit.Amount,
			&credit.InterestRate,
			&credit.Term,
			&credit.MonthlyPayment,
			&credit.RemainingAmount,
			&credit.Status,
			&credit.StartDate,
			&credit.EndDate,
			&credit.CreatedAt,
			&credit.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		credits = append(credits, credit)
	}
	return credits, nil
}
