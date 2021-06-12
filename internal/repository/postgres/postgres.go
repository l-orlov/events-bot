package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/l-orlov/events-bot/internal/config"
	_ "github.com/lib/pq"
)

func ConnectToDB(cfg config.PostgresDB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", initConnectionString(cfg))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func initConnectionString(cfg config.PostgresDB) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Address.Host, cfg.Address.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}
