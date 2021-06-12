package service

import (
	"github.com/l-orlov/events-bot/internal/repository"
)

type (
	UserProfile interface{}
	Service     struct {
		UserProfile
	}
)

func NewService(
	repo *repository.Repository,
) (*Service, error) {
	return &Service{
		UserProfile: NewUserProfileService(repo),
	}, nil
}
