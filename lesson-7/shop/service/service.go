package service

import (
	"context"
	"errors"
	"log"

	"shop/logger"
	"shop/models"
	tg "shop/pkg/tgbot"
	rep "shop/repository"
)

type Service interface {
	CreateItem(ctx context.Context, item *models.Item) (*models.Item, error)
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
}

type service struct {
	tg tg.TelegramAPI
	db rep.Repository
}

func (s *service) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	for _, itemID := range order.ItemIDs {
		_, err := s.db.GetItem(ctx, int32(itemID))
		if err != nil {
			return nil, errors.New("item not found")
		}
	}

	order, err := s.db.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	if err := s.tg.SendOrderNotification(order); err != nil {
		log.Println(err)
	}
	logger.Logger(ctx).Info("new order created")
	return order, err
}

func (s *service) CreateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	if item.Name == "" {
		return nil, errors.New("item name is empty")
	}
	if item.Price <= 0 {
		return nil, errors.New("item price should be positive")
	}
	logger.Logger(ctx).Error("new item created")
	return s.db.CreateItem(ctx, item)
}

func NewService(tg tg.TelegramAPI, db rep.Repository) Service {
	return &service{
		db: db,
		tg: tg,
	}
}
