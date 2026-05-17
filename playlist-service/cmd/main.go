package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bakdaulet212/ap2-final-/playlist-service/internal/handler"
	"github.com/bakdaulet212/ap2-final-/playlist-service/internal/repository"
	"github.com/bakdaulet212/ap2-final-/proto/playlistpb"
	"google.golang.org/grpc"
)

func main() {
	repo := repository.NewPostgresRepository()

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	playlistHandler := handler.NewPlaylistHandler(repo)
	playlistpb.RegisterPlaylistServiceServer(s, playlistHandler)

	fmt.Println("Playlist Service is running on port :50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}