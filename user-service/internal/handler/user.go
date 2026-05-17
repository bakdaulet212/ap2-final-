package handler

import (
	"context"
	"errors"

	"github.com/bakdaulet212/ap2-final-/proto/userpb"
	"github.com/bakdaulet212/ap2-final-/user-service/internal/models"
	"github.com/bakdaulet212/ap2-final-/user-service/internal/repository"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" || req.GetUsername() == "" {
		return nil, errors.New("missing required fields")
	}

	user := &models.User{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	id, err := h.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &userpb.RegisterResponse{
		UserId: id,
	}, nil
}