package config

import (
	cr "github.com/l-orlov/task-tracker/pkg/configreader"
)

type (
	Config struct {
		TelegramToken string     `yaml:"telegramToken" env:"TELEGRAM_TOKEN,default=test"`
		Logger        Logger     `yaml:"logger"`
		PostgresDB    PostgresDB `yaml:"postgresDB"`
	}
	Logger struct {
		Level  string `yaml:"level" env:"LOGGER_LEVEL,default=info"`
		Format string `yaml:"format" env:"LOGGER_FORMAT,default=default"`
	}
	PostgresDB struct {
		Address         cr.AddressConfig  `yaml:"address" env:"PG_ADDRESS,default=0.0.0.0:5432"`
		User            string            `yaml:"user" env:"PG_USER,default=postgres"`
		Password        string            `yaml:"password" env:"PG_PASSWORD,default=123"`
		Database        string            `yaml:"name" env:"PG_DATABASE,default=postgres"`
		SSLMode         string            `yaml:"sslmode" env:"PG_SSL_MODE,default=disable"`
		ConnMaxLifetime cr.DurationConfig `yaml:"connMaxLifetime"`
		MaxOpenConns    int               `yaml:"maxOpenConns"`
		MaxIdleConns    int               `yaml:"maxIdleConns"`
		Timeout         cr.DurationConfig `yaml:"timeout"`
		MigrationMode   bool              `yaml:"migrationMode"`
		MigrationDir    string            `yaml:"migrationDir"`
	}
)

func Init(path string) (*Config, error) {
	var cfg Config
	if err := cr.ReadYamlAndSetEnv(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
