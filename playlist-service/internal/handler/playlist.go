package handler

import (
	"context"
	"errors"

	"github.com/bakdaulet212/ap2-final-/playlist-service/internal/models"
	"github.com/bakdaulet212/ap2-final-/playlist-service/internal/repository"
	"github.com/bakdaulet212/ap2-final-/proto/playlistpb"
)

type PlaylistHandler struct {
	playlistpb.UnimplementedPlaylistServiceServer
	repo *repository.PlaylistRepository
}

func NewPlaylistHandler(repo *repository.PlaylistRepository) *PlaylistHandler {
	return &PlaylistHandler{repo: repo}
}

func (h *PlaylistHandler) CreatePlaylist(ctx context.Context, req *playlistpb.CreatePlaylistRequest) (*playlistpb.CreatePlaylistResponse, error) {
	if req.GetName() == "" || req.GetUserId() == "" {
		return nil, errors.New("name and user_id are required")
	}

	playlist := &models.Playlist{
		Title:  req.GetName(),
		UserID: req.GetUserId(),
	}

	id, err := h.repo.CreatePlaylist(playlist)
	if err != nil {
		return nil, err
	}

	return &playlistpb.CreatePlaylistResponse{
		Playlist: &playlistpb.Playlist{
			Id: id,
		},
	}, nil
}