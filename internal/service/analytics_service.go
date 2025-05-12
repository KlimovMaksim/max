package service

import (
	"context"
	"max/internal/repository"
)

type AnalyticsService struct {
	transactionRepo *repository.TransactionRepository
	creditRepo      *repository.CreditRepository
	accountRepo     *repository.AccountRepository
}

func NewAnalyticsService(transactionRepo *repository.TransactionRepository, creditRepo *repository.CreditRepository, accountRepo *repository.AccountRepository) *AnalyticsService {
	return &AnalyticsService{
		transactionRepo: transactionRepo,
		creditRepo:      creditRepo,
		accountRepo:     accountRepo,
	}
}

// Пример: получение общей аналитики по пользователю
func (s *AnalyticsService) GetAnalytics(ctx context.Context, userID int64) (interface{}, error) {
	// Получаем счета пользователя
	accounts, err := s.accountRepo.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	var totalBalance float64
	var totalCredits float64
	var totalTransactions int
	for _, acc := range accounts {
		totalBalance += acc.Balance
		// Получаем кредиты по каждому счету
		credits, err := s.creditRepo.ListByAccount(ctx, acc.ID)
		if err == nil {
			for _, c := range credits {
				totalCredits += c.Amount
			}
		}
		// Получаем транзакции по каждому счету
		transactions, err := s.transactionRepo.ListByAccount(ctx, acc.ID)
		if err == nil {
			totalTransactions += len(transactions)
		}
	}
	return map[string]interface{}{
		"total_balance":      totalBalance,
		"total_credits":      totalCredits,
		"total_transactions": totalTransactions,
	}, nil
}

// Пример: получение кредитной нагрузки пользователя
func (s *AnalyticsService) GetCreditLoad(ctx context.Context, userID int64) (interface{}, error) {
	accounts, err := s.accountRepo.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	var totalCredits float64
	var totalBalance float64
	for _, acc := range accounts {
		credits, err := s.creditRepo.ListByAccount(ctx, acc.ID)
		if err == nil {
			for _, c := range credits {
				totalCredits += c.Amount
			}
		}
		totalBalance += acc.Balance
	}
	var creditLoad float64
	if totalBalance > 0 {
		creditLoad = totalCredits / totalBalance
	} else {
		creditLoad = 0
	}
	return map[string]interface{}{
		"credit_load": creditLoad,
	}, nil
}
