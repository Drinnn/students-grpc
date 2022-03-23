package main

import (
	"log"
	"net"

	"github.com/Drinnn/students-grpc/protos"
	"github.com/Drinnn/students-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	protos.RegisterStudentServiceServer(grpcServer, services.NewStudentService())

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}
