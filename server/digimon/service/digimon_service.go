package service

import (
	"context"
	"errors"
	"server/domain"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type digimonService struct {
	digimonRepo domain.DigimonRepository
}

func NewDigimonService(digimonRepo domain.DigimonRepository) domain.DigimonService {
	return &digimonService{
		digimonRepo: digimonRepo,
	}
}

// Implement DigimonService

func (ds *digimonService) GetById(ctx context.Context, id string) (*domain.Digimon, error) {
	digimon, err := ds.digimonRepo.GetById(ctx, id)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return digimon, nil
}

func (ds *digimonService) Store(ctx context.Context, d *domain.Digimon) error {
	if d.Id == "" {
		d.Id = uuid.Must(uuid.NewV4()).String()
	}
	if d.Status == "" {
		d.Status = "good"
	}
	if err := ds.digimonRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (ds *digimonService) UpdateStatus(ctx context.Context, d *domain.Digimon) error {
	if d.Status == "" {
		err := errors.New("status is blank")
		logrus.Error(err)
		return err
	}

	if err := ds.digimonRepo.UpdateStatus(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
