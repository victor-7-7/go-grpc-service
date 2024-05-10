package grpc

import (
	"context"
	"go-course/grpc-service/internal/rocket"
	rkt "go-course/grpc-service/proto/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Для уменьшения связанности и улучшения тестирования не будем
// импортировать зависимость rocket.Service, а обойдемся интерфейсом.

// RocketService - defines the interface that the concrete
// implementation has to adhere to
type RocketService interface {
	GetRocketById(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	rkt.UnimplementedRocketServiceServer
	RocketService RocketService
}

// New - returns a new gRPC handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("could not listen on :50051")
		return err
	}
	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve %s\n", err)
		return err
	}
	return nil
}

// GetRocket -retrieves a rocket by id and returns response.
func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Println("Handler | Get Rocket gRPC Endpoint Hit.")
	rocketObj, err := h.RocketService.GetRocketById(ctx, req.GetId())
	if err != nil {
		log.Println("Failed to retrieve rocket by id.")
		return &rkt.GetRocketResponse{}, err
	}
	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocketObj.Id,
			Name: rocketObj.Name,
			Type: rocketObj.Type,
		},
	}, nil
}

// AddRocket -adds a rocket to the database.
func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Println("Add Rocket gRPC Endpoint Hit.")
	rocketFromDb, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		Id:   req.Rocket.Id,
		Name: req.Rocket.Name,
		Type: req.Rocket.Type,
	})
	if err != nil {
		log.Println("Failed to insert rocket into database.")
		return &rkt.AddRocketResponse{}, err
	}
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocketFromDb.Id,
			Name: rocketFromDb.Name,
			Type: rocketFromDb.Type,
		},
	}, nil
}

// DeleteRocket - handler for deleting a rocket.
func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Println("Delete Rocket gRPC Endpoint Hit.")
	err := h.RocketService.DeleteRocket(ctx, req.Rocket.GetId())
	if err != nil {
		log.Println("Failed to delete rocket.")
		return &rkt.DeleteRocketResponse{}, err
	}
	return &rkt.DeleteRocketResponse{
		Status: "successfully deleted rocket",
	}, nil
}
