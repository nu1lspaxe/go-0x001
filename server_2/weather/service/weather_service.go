package service

import (
	"context"

	"server_2/domain"

	"github.com/sirupsen/logrus"
)

type weatherService struct {
	weatherRepo domain.WeatherRepository
}

func NewWeatherService(weatherRepo domain.WeatherRepository) domain.WeatherService {
	return &weatherService{
		weatherRepo: weatherRepo,
	}
}

func (w *weatherService) GetByLocation(ctx context.Context, location string) (*domain.Weather, error) {
	weather, err := w.weatherRepo.GetByLocation(ctx, location)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return weather, nil
}
