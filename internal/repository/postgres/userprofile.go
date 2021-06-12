package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	userProfileTable = "userprofile"
)

type UserProfilePostgres struct {
	db        *sqlx.DB
	dbTimeout time.Duration
}

func NewUserProfilePostgres(db *sqlx.DB, dbTimeout time.Duration) *UserProfilePostgres {
	return &UserProfilePostgres{
		db:        db,
		dbTimeout: dbTimeout,
	}
}
