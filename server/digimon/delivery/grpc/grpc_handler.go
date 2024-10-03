package grpc

import (
	"context"
	"time"

	"github.com/nu1lspaxe/go-0x001/server/domain"
	pb "github.com/nu1lspaxe/go-0x001/server/proto/digimon"

	"github.com/sirupsen/logrus"
	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DigimonHandler ...
type DigimonHandler struct {
	DigimonServ domain.DigimonService
	DietServ    domain.DietService
	WeatherServ domain.WeatherService
	pb.UnimplementedDigimonServer
}

// NewDigimonHandler ...
func NewDigimonHandler(s *grpcLib.Server, digimonServ domain.DigimonService, dietServ domain.DietService, weatherServ domain.WeatherService) {
	handler := &DigimonHandler{
		DigimonServ: digimonServ,
		DietServ:    dietServ,
		WeatherServ: weatherServ,
	}

	pb.RegisterDigimonServer(s, handler)
}

func (d *DigimonHandler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	digimon := domain.Digimon{
		Name: req.GetName(),
	}
	if err := d.DigimonServ.Store(ctx, &digimon); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	return &pb.CreateResponse{
		Id:     digimon.Id,
		Name:   digimon.Name,
		Status: digimon.Status,
	}, nil
}

func (d *DigimonHandler) QueryStream(req *pb.QueryRequest, srv pb.Digimon_QueryStreamServer) error {
	weatherClient, err := d.WeatherServ.GetStreamByLocation(context.Background(), "A")
	if err != nil {
		logrus.Error(err)
		return err
	}

	for {
		if err := weatherClient.Send(&domain.Weather{
			Location: "A",
		}); err != nil {
			logrus.Error(err)
			return err
		}

		time.Sleep(time.Duration(5) * time.Second)

		weather, err := weatherClient.Recv()
		if err != nil {
			logrus.Error(err)
			return err
		}

		digimon, err := d.DigimonServ.GetById(context.Background(), req.GetId())
		if err != nil {
			logrus.Error(err)
			return err
		}

		srv.Send(&pb.QueryResponse{
			Id:       digimon.Id,
			Name:     digimon.Name,
			Status:   digimon.Status,
			Location: weather.Location,
			Weather:  weather.Weather,
		})
	}
}

func (d *DigimonHandler) Foster(ctx context.Context, req *pb.FosterRequest) (*pb.FosterResponse, error) {
	if err := d.DietServ.Store(ctx, &domain.Diet{
		UserId: req.GetId(),
		Name:   req.GetFood().GetName(),
	}); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	if err := d.DigimonServ.UpdateStatus(ctx, &domain.Digimon{
		Id:     req.GetId(),
		Status: "good",
	}); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	return &pb.FosterResponse{}, nil
}
