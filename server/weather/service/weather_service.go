package service

import (
	"context"

	"github.com/nu1lspaxe/go-0x001/server/domain"

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

func (wu *weatherService) GetStreamByLocation(ctx context.Context, location string) (domain.StreamWeather, error) {
	client, err := wu.weatherRepo.GetStreamByLocation(ctx, location)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return client, nil
}
