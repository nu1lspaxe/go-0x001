package domain

import (
	"context"
)

type WeatherEnum int32

const (
	SUNNY  WeatherEnum = 0
	CLOUDY WeatherEnum = 1
)

type Weather struct {
	Location string
	Weather  WeatherEnum
}

type WeatherRepository interface {
	GetByLocation(ctx context.Context, location string) (*Weather, error)
}

type WeatherService interface {
	GetByLocation(ctx context.Context, location string) (*Weather, error)
}
