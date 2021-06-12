package service

import (
	"github.com/l-orlov/events-bot/internal/repository"
)

type (
	UserProfileService struct {
		repo *repository.Repository
	}
)

func NewUserProfileService(repo *repository.Repository) *UserProfileService {
	return &UserProfileService{
		repo: repo,
	}
}
