package service

import (
	"context"

	"github.com/nu1lspaxe/go-0x001/server/domain"

	"github.com/sirupsen/logrus"
)

type dietService struct {
	dietRepo domain.DietRepository
}

func NewDietService(dietRepo domain.DietRepository) domain.DietService {
	return &dietService{
		dietRepo,
	}
}

func (ds *dietService) GetById(ctx context.Context, id string) (*domain.Diet, error) {
	diet, err := ds.dietRepo.GetById(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return diet, nil
}

func (ds *dietService) Store(ctx context.Context, d *domain.Diet) error {
	if err := ds.dietRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
