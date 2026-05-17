package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bakdaulet212/ap2-final-/user-service/internal/handler"
	"github.com/bakdaulet212/ap2-final-/user-service/internal/repository"
	"github.com/bakdaulet212/ap2-final-/proto/userpb"
	"google.golang.org/grpc"
)

func main() {
	repo := repository.NewPostgresRepository()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userHandler := handler.NewUserHandler(repo)
	
	userpb.RegisterUserServiceServer(s, userHandler)

	fmt.Println("User Service is running on port :50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}