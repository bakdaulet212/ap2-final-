package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/bakdaulet212/ap2-final-/proto/catalogpb"
	"github.com/bakdaulet212/ap2-final-/proto/playlistpb"
	"github.com/bakdaulet212/ap2-final-/proto/userpb"
	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	userClient     userpb.UserServiceClient
	catalogClient  catalogpb.CatalogServiceClient
	playlistClient playlistpb.PlaylistServiceClient
}

func NewGatewayHandler(u userpb.UserServiceClient, c catalogpb.CatalogServiceClient, p playlistpb.PlaylistServiceClient) *GatewayHandler {
	return &GatewayHandler{
		userClient:     u,
		catalogClient:  c,
		playlistClient: p,
	}
}

func (h *GatewayHandler) RegisterUser(c *gin.Context) {
	var req userpb.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := h.userClient.Register(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *GatewayHandler) AddTrack(c *gin.Context) {
	var req catalogpb.AddTrackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := h.catalogClient.AddTrack(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *GatewayHandler) GetTrack(c *gin.Context) {
	trackID := c.Param("id")
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := h.catalogClient.GetTrack(ctx, &catalogpb.GetTrackRequest{TrackId: trackID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *GatewayHandler) CreatePlaylist(c *gin.Context) {
	var req playlistpb.CreatePlaylistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := h.playlistClient.CreatePlaylist(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}