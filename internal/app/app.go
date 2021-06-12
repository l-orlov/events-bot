package app

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/l-orlov/events-bot/internal/config"
	"github.com/l-orlov/events-bot/internal/repository"
	"github.com/l-orlov/events-bot/internal/repository/postgres"
	"github.com/l-orlov/events-bot/internal/service"
	"github.com/l-orlov/events-bot/internal/telegram"
	"github.com/l-orlov/task-tracker/pkg/logger"
)

// Run initializes whole application.
func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	lg, err := logger.New(cfg.Logger.Level, cfg.Logger.Format)
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	// Dependencies
	db, err := postgres.ConnectToDB(cfg.PostgresDB)
	if err != nil {
		lg.Fatalf("failed to connect to db: %v", err)
	}
	defer func() {
		if err = db.Close(); err != nil {
			lg.Errorf("failed to close db: %v", err)
		}
	}()

	if cfg.PostgresDB.MigrationMode {
		if err = postgres.MigrateSchema(db.DB, cfg.PostgresDB); err != nil {
			lg.Fatalf("failed to do migration: %v", err)
		}
	}

	// Repo, Service & API Handlers
	repo, err := repository.NewRepository(cfg, db)
	if err != nil {
		lg.Fatalf("failed to create repository: %v", err)
	}

	svc, err := service.NewService(repo)
	if err != nil {
		lg.Fatalf("failed to create service: %v", err)
	}

	_ = svc

	// Telegram bot
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		lg.Fatal(err)
	}

	bot := telegram.NewBot(botApi)
	if err := bot.Start(); err != nil {
		lg.Fatal(err)
	}
}
