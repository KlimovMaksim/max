package service

import (
	"max/internal/models"
	"max/internal/repository"
	"context"
	"errors"
)

type CreditService struct {
	repo *repository.CreditRepository
}

func NewCreditService(repo *repository.CreditRepository) *CreditService {
	return &CreditService{repo: repo}
}

func (s *CreditService) CreateCredit(ctx context.Context, credit *models.Credit) error {
	if credit.Amount <= 0 {
		return errors.New("сумма кредита должна быть положительной")
	}
	s.repo.Create(ctx, credit)
	return nil
}

func (s *CreditService) GetCreditByID(ctx context.Context, id int64) (*models.Credit, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CreditService) ListCreditsByAccount(ctx context.Context, accountID int64) ([]models.Credit, error) {
	return s.repo.ListByAccount(ctx, accountID)
}
