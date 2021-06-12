package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/l-orlov/events-bot/internal/config"
	"github.com/l-orlov/events-bot/internal/repository/postgres"
)

type (
	UserProfile interface{}
	Repository  struct {
		UserProfile
	}
)

func NewRepository(
	cfg *config.Config, db *sqlx.DB,
) (*Repository, error) {
	userProfileRepo := postgres.NewUserProfilePostgres(db, cfg.PostgresDB.Timeout.Duration())

	return &Repository{
		UserProfile: userProfileRepo,
	}, nil
}
