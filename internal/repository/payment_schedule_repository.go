package repository

import (
	"max/internal/models"
	"context"
	"database/sql"
	"errors"
	"time"
)

type PaymentScheduleRepository struct {
	db *sql.DB
}

func NewPaymentScheduleRepository(db *sql.DB) *PaymentScheduleRepository {
	return &PaymentScheduleRepository{db: db}
}

func (r *PaymentScheduleRepository) Create(ctx context.Context, ps *models.PaymentSchedule) error {
	query := `INSERT INTO payment_schedules (credit_id, account_id, due_date, amount, penalty, status, paid_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query,
		ps.CreditID,
		ps.AccountID,
		ps.DueDate,
		ps.Amount,
		ps.Penalty,
		ps.Status,
		ps.PaidAt,
		time.Now(),
		time.Now(),
	)
	return err
}

func (r *PaymentScheduleRepository) GetByID(ctx context.Context, id int64) (*models.PaymentSchedule, error) {
	query := `SELECT id, credit_id, account_id, due_date, amount, penalty, status, paid_at, created_at, updated_at FROM payment_schedules WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)
	var ps models.PaymentSchedule
	err := row.Scan(
		&ps.ID,
		&ps.CreditID,
		&ps.AccountID,
		&ps.DueDate,
		&ps.Amount,
		&ps.Penalty,
		&ps.Status,
		&ps.PaidAt,
		&ps.CreatedAt,
		&ps.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("платеж не найден")
	}
	return &ps, err
}

func (r *PaymentScheduleRepository) ListByCredit(ctx context.Context, creditID int64) ([]models.PaymentSchedule, error) {
	query := `SELECT id, credit_id, account_id, due_date, amount, penalty, status, paid_at, created_at, updated_at FROM payment_schedules WHERE credit_id = ?`
	rows, err := r.db.QueryContext(ctx, query, creditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var schedules []models.PaymentSchedule
	for rows.Next() {
		var ps models.PaymentSchedule
		err := rows.Scan(
			&ps.ID,
			&ps.CreditID,
			&ps.AccountID,
			&ps.DueDate,
			&ps.Amount,
			&ps.Penalty,
			&ps.Status,
			&ps.PaidAt,
			&ps.CreatedAt,
			&ps.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, ps)
	}
	return schedules, nil
}
