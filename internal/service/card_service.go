package service

import (
	"context"
	"errors"
	"max/internal/models"
	"max/internal/repository"
	"time"
)

type CardService struct {
	repo *repository.CardRepository
}

func NewCardService(repo *repository.CardRepository) *CardService {
	return &CardService{repo: repo}
}

func (s *CardService) CreateCard(ctx context.Context, card *models.Card) error {
	if card.Number == "" || card.Expiry.IsZero() || len(card.CVV) == 0 {
		return errors.New("некорректные данные карты")
	}
	card.CreatedAt = time.Now()
	return s.repo.Create(ctx, card)
}

func (s *CardService) GetCardByID(ctx context.Context, id int64) (*models.Card, error) {
	card, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *CardService) ListCardsByAccount(ctx context.Context, accountID int64) ([]models.Card, error) {
	cards, err := s.repo.ListByAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return cards, nil
}
