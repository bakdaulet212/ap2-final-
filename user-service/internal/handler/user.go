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
	if req.GetUsername() == "" || req.GetEmail() == "" || req.GetPassword() == "" {
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
		UserId:  id,
		Message: "User registered successfully!",
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	return &userpb.LoginResponse{
		Message: "Login successful (stub)",
	}, nil
}

func (h *UserHandler) GetProfile(ctx context.Context, req *userpb.GetProfileRequest) (*userpb.GetProfileResponse, error) {
	return &userpb.GetProfileResponse{
		Id:       req.GetUserId(),
		Username: "testuser",
		Email:    "test@example.com",
	}, nil
}

func (h *UserHandler) UpdateProfile(ctx context.Context, req *userpb.UpdateProfileRequest) (*userpb.UpdateProfileResponse, error) {
	return &userpb.UpdateProfileResponse{
		Success: true,
		Message: "Profile updated successfully",
	}, nil
}