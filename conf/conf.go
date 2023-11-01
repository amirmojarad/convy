package conf

import (
	"os"
	"strconv"
)

type AppConfig struct {
	Database struct {
		Name     string
		Port     int
		Host     string
		Password string
		Username string
	}

	App struct {
		Debug bool
		Name  string
		Port  int
	}
}

func NewAppConfig() (*AppConfig, error) {
	var cfg AppConfig

	if err := setDatabase(&cfg); err != nil {
		return nil, err
	}

	if err := setApp(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setDatabase(cfg *AppConfig) error {
	port, err := strconv.Atoi(os.Getenv("DATABASE_POSTGRES_PORT"))
	if err != nil {
		return err
	}

	cfg.Database.Port = port
	cfg.Database.Name = os.Getenv("DATABASE_POSTGRES_NAME")
	cfg.Database.Host = os.Getenv("DATABASE_POSTGRES_HOST")
	cfg.Database.Password = os.Getenv("DATABASE_POSTGRES_PASSWORD")
	cfg.Database.Username = os.Getenv("DATABASE_POSTGRES_USER")

	return nil
}

func setApp(cfg *AppConfig) error {
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return err
	}

	cfg.App.Debug = os.Getenv("APP_DEBUG") == "true"
	cfg.App.Port = port
	cfg.App.Name = os.Getenv("APP_NAME")

	return nil
}
