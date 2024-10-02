package domain

import "context"

type Weather struct {
	Location string
	Weather  string
}

type StreamWeather interface {
	Send(*Weather) error
	Recv() (*Weather, error)
}

type WeatherRepository interface {
	GetStreamByLocation(ctx context.Context, location string) (StreamWeather, error)
}

type WeatherService interface {
	GetStreamByLocation(ctx context.Context, location string) (StreamWeather, error)
}
