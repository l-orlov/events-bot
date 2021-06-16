package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// Telegram bot
	bot, err := telegram.NewBot(lg, svc, cfg.TelegramToken)
	go func() {
		bot.Start()
	}()
	lg.Infof("bot started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	lg.Info("bot shutting down")
	bot.Stop()
}
