package service

import (
	"max/internal/models"
	"max/internal/repository"
	"context"
	"errors"
	"time"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	if transaction.Amount <= 0 {
		return errors.New("сумма транзакции должна быть положительной")
	}
	if transaction.Type == "" {
		return errors.New("тип транзакции не указан")
	}
	transaction.CreatedAt = time.Now()
	return s.repo.Create(ctx, transaction)
}

func (s *TransactionService) GetTransactionByID(ctx context.Context, id int64) (*models.Transaction, error) {
	transaction, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) ListTransactionsByAccount(ctx context.Context, accountID int64) ([]models.Transaction, error) {
	transactions, err := s.repo.ListByAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
