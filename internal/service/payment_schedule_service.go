package service

import (
	"max/internal/models"
	"max/internal/repository"
	"context"
	"errors"
)

type PaymentScheduleService struct {
	repo *repository.PaymentScheduleRepository
}

func NewPaymentScheduleService(repo *repository.PaymentScheduleRepository) *PaymentScheduleService {
	return &PaymentScheduleService{repo: repo}
}

func (s *PaymentScheduleService) CreatePayment(ctx context.Context, ps *models.PaymentSchedule) error {
	if ps.Amount <= 0 {
		return errors.New("сумма платежа должна быть положительной")
	}
	return s.repo.Create(ctx, ps)
}

func (s *PaymentScheduleService) GetPaymentByID(ctx context.Context, id int64) (*models.PaymentSchedule, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PaymentScheduleService) ListPaymentsByCredit(ctx context.Context, creditID int64) ([]models.PaymentSchedule, error) {
	return s.repo.ListByCredit(ctx, creditID)
}
