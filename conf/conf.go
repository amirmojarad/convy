package conf

import (
	"fmt"
	_ "gorm.io/driver/postgres"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	Database struct {
		Name                  string
		Port                  int
		Host                  string
		Password              string
		Username              string
		SslMode               string
		Timezone              string
		MaxIdleConnections    int
		MaxOpenConnections    int
		ConnectionMaxLifetime time.Duration
		MigrationPath         string
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
	cfg.Database.Username = os.Getenv("DATABASE_POSTGRES_USERNAME")
	cfg.Database.Password = os.Getenv("DATABASE_POSTGRES_PASSWORD")
	cfg.Database.Host = os.Getenv("DATABASE_POSTGRES_HOST")
	cfg.Database.Name = os.Getenv("DATABASE_POSTGRES_NAME")

	port, err := envConvertor("DATABASE_POSTGRES_PORT", func(v string) (uint64, error) {
		return strconv.ParseUint(v, 10, 32)
	})
	if err != nil {
		return err
	}

	cfg.Database.Port = int(port)

	cfg.Database.SslMode = os.Getenv("DATABASE_POSTGRES_SSLMODE")
	cfg.Database.Timezone = os.Getenv("DATABASE_POSTGRES_TIMEZONE")

	maxConn, err := envConvertor("DATABASE_POSTGRES_MAX_OPEN_CONN", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.Database.MaxOpenConnections = maxConn

	maxIdle, err := envConvertor("DATABASE_POSTGRES_MAX_IDLE_CONN", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.Database.MaxIdleConnections = maxIdle

	connMaxLif, err := envConvertor("DATABASE_POSTGRES_CONN_MAX_LIFETIME", time.ParseDuration)
	if err != nil {
		return err
	}

	cfg.Database.ConnectionMaxLifetime = connMaxLif

	cfg.Database.MigrationPath = os.Getenv("DATABASE_POSTGRES_MIGRATION_PATH")

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

func envConvertor[T any](envKey string, converter func(v string) (T, error)) (T, error) {
	value := os.Getenv(envKey)

	result, err := converter(value)
	if err != nil {
		var noop T

		return noop, fmt.Errorf("%s is not a valid value for %s, %w", value, envKey, err)
	}

	return result, nil
}
