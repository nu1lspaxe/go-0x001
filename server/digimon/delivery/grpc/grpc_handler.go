package grpc

import (
	"context"

	"server/domain"
	pb "server/proto"

	"github.com/sirupsen/logrus"
	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DigimonHandler struct {
	DigimonServ domain.DigimonService
	DietServ    domain.DietService
	pb.UnimplementedDigimonServer
}

func NewDigimonHandler(s *grpcLib.Server, digimonServ domain.DigimonService, dietServ domain.DietService) {
	handler := &DigimonHandler{
		DigimonServ: digimonServ,
		DietServ:    dietServ,
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

func (d *DigimonHandler) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	digimon, err := d.DigimonServ.GetById(ctx, req.GetId())
	if err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Query digimon error")
	}

	return &pb.QueryResponse{
		Id:     digimon.Id,
		Name:   digimon.Name,
		Status: digimon.Status,
	}, nil
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
