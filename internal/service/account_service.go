package service

import (
	"max/internal/models"
	"max/internal/repository"
	"context"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(ctx context.Context, account *models.Account) error {
	// TODO: реализовать бизнес-логику создания счета
	return s.repo.Create(ctx, account)
}

func (s *AccountService) GetAccountByID(ctx context.Context, id int64) (*models.Account, error) {
	// TODO: реализовать бизнес-логику получения счета по ID
	return s.repo.GetByID(ctx, id)
}

func (s *AccountService) ListAccountsByUser(ctx context.Context, userID int64) ([]models.Account, error) {
	// TODO: реализовать бизнес-логику получения счетов пользователя
	return s.repo.ListByUser(ctx, userID)
}
