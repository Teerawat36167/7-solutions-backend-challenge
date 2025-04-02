package api

import (
	pb "github.com/Teerawat36167/PieFireDire/internal/pb"
	counter "github.com/Teerawat36167/PieFireDire/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedBeefServiceServer
	counter *counter.MeatCounter
}

func NewServer() *Server {
	return &Server{
		counter: counter.NewMeatCounter(),
	}
}

func StartGRPCServer(addr string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	beefService := NewServer()

	pb.RegisterBeefServiceServer(grpcServer, beefService)

	reflection.Register(grpcServer)

	return grpcServer, nil
}
