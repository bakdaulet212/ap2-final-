package main

import (
	"log"

	"github.com/bakdaulet212/ap2-final-/api-gateway/internal/handler"
	"github.com/bakdaulet212/ap2-final-/proto/catalogpb"
	"github.com/bakdaulet212/ap2-final-/proto/playlistpb"
	"github.com/bakdaulet212/ap2-final-/proto/userpb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	userConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	catalogConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to catalog service: %v", err)
	}
	defer catalogConn.Close()

	playlistConn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to playlist service: %v", err)
	}
	defer playlistConn.Close()

	userClient := userpb.NewUserServiceClient(userConn)
	catalogClient := catalogpb.NewCatalogServiceClient(catalogConn)
	playlistClient := playlistpb.NewPlaylistServiceClient(playlistConn)

	gwHandler := handler.NewGatewayHandler(userClient, catalogClient, playlistClient)

	r := gin.Default()

	r.POST("/register", gwHandler.RegisterUser)
	r.POST("/tracks", gwHandler.AddTrack)
	r.GET("/tracks/:id", gwHandler.GetTrack)
	r.POST("/playlists", gwHandler.CreatePlaylist)

	log.Println("API Gateway is running on port :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run gateway: %v", err)
	}
}