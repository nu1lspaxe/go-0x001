package fake

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/nu1lspaxe/go-0x001/server_2/domain"
)

type fakeWeatherRepository struct {
}

func NewFakeWeatherRepository() domain.WeatherRepository {
	return &fakeWeatherRepository{}
}

func createRandomWeather() domain.WeatherEnum {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	switch randomNum := rand.Intn(2); randomNum {
	case 0:
		return domain.SUNNY
	case 1:
		return domain.CLOUDY
	default:
		return domain.SUNNY
	}
}

func (f *fakeWeatherRepository) GetByLocation(ctx context.Context, location string) (*domain.Weather, error) {
	switch location {
	case "A":
		return &domain.Weather{
			Location: "A",
			Weather:  createRandomWeather(),
		}, nil
	default:
		return nil, errors.New("this location does not exist")
	}
}
