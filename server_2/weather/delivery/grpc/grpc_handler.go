package grpc

import (
	"context"
	"errors"
	"io"

	"github.com/nu1lspaxe/go-0x001/server_2/domain"
	pb "github.com/nu1lspaxe/go-0x001/server_2/proto/weather"

	"github.com/sirupsen/logrus"
	grpcLib "google.golang.org/grpc"
)

type WeatherHandler struct {
	WeatherServ domain.WeatherService
	pb.UnimplementedWeatherServer
}

func NewWeatherHandler(s *grpcLib.Server, weatherServ domain.WeatherService) {
	handler := &WeatherHandler{
		WeatherServ: weatherServ,
	}

	pb.RegisterWeatherServer(s, handler)
}

func mappingGrpcWeatherEnum(weather domain.WeatherEnum) (pb.QueryResponse_Weather, error) {
	switch weather {
	case domain.SUNNY:
		return pb.QueryResponse_SUNNY, nil
	case domain.CLOUDY:
		return pb.QueryResponse_CLOUDY, nil
	default:
		return pb.QueryResponse_SUNNY, errors.New("this weather does not exist")
	}
}

func (w *WeatherHandler) Query(srv pb.Weather_QueryServer) error {
	for {
		msg, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			logrus.Error(err)
			return err
		}

		weather, err := w.WeatherServ.GetByLocation(context.Background(), msg.GetLocation())
		if err != nil {
			logrus.Error(err)
		}

		grpcWeatherEnum, err := mappingGrpcWeatherEnum(weather.Weather)
		if err != nil {
			logrus.Error(err)
		}

		srv.Send(&pb.QueryResponse{
			Location: weather.Location,
			Weather:  grpcWeatherEnum,
		})
	}
}
