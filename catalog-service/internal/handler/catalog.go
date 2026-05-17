package handler

import (
	"context"
	"errors"

	"github.com/bakdaulet212/ap2-final-/catalog-service/internal/models"
	"github.com/bakdaulet212/ap2-final-/catalog-service/internal/repository"
	"github.com/bakdaulet212/ap2-final-/proto/catalogpb"
)

type CatalogHandler struct {
	catalogpb.UnimplementedCatalogServiceServer
	repo *repository.CatalogRepository
}

func NewCatalogHandler(repo *repository.CatalogRepository) *CatalogHandler {
	return &CatalogHandler{repo: repo}
}

func (h *CatalogHandler) AddTrack(ctx context.Context, req *catalogpb.AddTrackRequest) (*catalogpb.AddTrackResponse, error) {
	if req.GetTitle() == "" || req.GetArtist() == "" {
		return nil, errors.New("title and artist are required")
	}

	track := &models.Track{
		Title:    req.GetTitle(),
		Artist:   req.GetArtist(),
		Duration: req.GetDuration(),
	}

	id, err := h.repo.AddTrack(track)
	if err != nil {
		return nil, err
	}

	return &catalogpb.AddTrackResponse{
		Track: &catalogpb.Track{
			Id:       id,
			Title:    track.Title,
			Artist:   track.Artist,
			Duration: track.Duration,
		},
	}, nil
}

func (h *CatalogHandler) GetTrack(ctx context.Context, req *catalogpb.GetTrackRequest) (*catalogpb.GetTrackResponse, error) {
	track, err := h.repo.GetTrack(req.GetTrackId())
	if err != nil {
		return nil, errors.New("track not found")
	}

	return &catalogpb.GetTrackResponse{
		Track: &catalogpb.Track{
			Id:       track.ID,
			Title:    track.Title,
			Artist:   track.Artist,
			Duration: track.Duration,
		},
	}, nil
}