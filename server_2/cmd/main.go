package main

import (
	"net"
	_weatherHandlerGrpcDelivery "server_2/weather/delivery/grpc"
	_weatherRepo "server_2/weather/repository/fake"
	_weatherService "server_2/weather/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Fatal error config file: %v\n", err)
	}
}

func main() {
	logrus.Info("GRPC server started")

	grpcPort := viper.GetString("GRPC_PORT")

	weatherRepo := _weatherRepo.NewFakeWeatherRepository()

	weatherService := _weatherService.NewWeatherService(weatherRepo)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		logrus.Fatal(err)
	}
	s := grpc.NewServer()

	_weatherHandlerGrpcDelivery.NewWeatherHandler(s, weatherService)

	logrus.Fatal(s.Serve(lis))
}
