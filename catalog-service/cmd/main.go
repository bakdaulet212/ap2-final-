package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bakdaulet212/ap2-final-/catalog-service/internal/handler"
	"github.com/bakdaulet212/ap2-final-/catalog-service/internal/repository"
	"github.com/bakdaulet212/ap2-final-/proto/catalogpb"
	"google.golang.org/grpc"
)

func main() {
	repo := repository.NewCatalogRepository()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	catalogHandler := handler.NewCatalogHandler(repo)
	catalogpb.RegisterCatalogServiceServer(s, catalogHandler)

	fmt.Println("Catalog Service is running on port :50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
