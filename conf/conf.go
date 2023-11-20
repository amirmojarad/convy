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
		Mongo struct {
			Username          string
			Password          string
			Host              string
			Port              string
			ConnectionOptions struct {
				ConnectionTimeout time.Duration
				MaxPoolSize       uint64
				W                 string
			}
		}
		Postgres struct {
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
	}

	App struct {
		Debug bool
		Name  string
		Port  int
	}

	Secrets struct {
		EncryptionKey []byte
	}

	TokenDetails struct {
		Issuer        string
		Secret        string
		AtExpiresDays int
		RtExpiresDays int
	}
}

func NewAppConfig() (*AppConfig, error) {
	var cfg AppConfig

	if err := setPostgres(&cfg); err != nil {
		return nil, err
	}

	if err := setApp(&cfg); err != nil {
		return nil, err
	}

	setSecrets(&cfg)

	if err := setTokenDetails(&cfg); err != nil {
		return nil, err
	}

	if err := setMongoDb(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setPostgres(cfg *AppConfig) error {
	cfg.Database.Postgres.Username = os.Getenv("DATABASE_POSTGRES_USERNAME")
	cfg.Database.Postgres.Password = os.Getenv("DATABASE_POSTGRES_PASSWORD")
	cfg.Database.Postgres.Host = os.Getenv("DATABASE_POSTGRES_HOST")
	cfg.Database.Postgres.Name = os.Getenv("DATABASE_POSTGRES_NAME")

	port, err := envConvertor("DATABASE_POSTGRES_PORT", func(v string) (uint64, error) {
		return strconv.ParseUint(v, 10, 32)
	})
	if err != nil {
		return err
	}

	cfg.Database.Postgres.Port = int(port)

	cfg.Database.Postgres.SslMode = os.Getenv("DATABASE_POSTGRES_SSLMODE")
	cfg.Database.Postgres.Timezone = os.Getenv("DATABASE_POSTGRES_TIMEZONE")

	maxConn, err := envConvertor("DATABASE_POSTGRES_MAX_OPEN_CONN", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.Database.Postgres.MaxOpenConnections = maxConn

	maxIdle, err := envConvertor("DATABASE_POSTGRES_MAX_IDLE_CONN", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.Database.Postgres.MaxIdleConnections = maxIdle

	connMaxLif, err := envConvertor("DATABASE_POSTGRES_CONN_MAX_LIFETIME", time.ParseDuration)
	if err != nil {
		return err
	}

	cfg.Database.Postgres.ConnectionMaxLifetime = connMaxLif

	cfg.Database.Postgres.MigrationPath = os.Getenv("DATABASE_POSTGRES_MIGRATION_PATH")

	return nil
}

func setMongoDb(cfg *AppConfig) error {
	cfg.Database.Mongo.Password = os.Getenv("DATABASE_MONGO_PASSWORD")
	cfg.Database.Mongo.Host = os.Getenv("DATABASE_MONGO_HOST")
	cfg.Database.Mongo.Username = os.Getenv("DATABASE_MONGO_USERNAME")
	cfg.Database.Mongo.Port = os.Getenv("DATABASE_MONGO_PORT")

	connectionTimeout, err := envConvertor("DATABASE_MONGO_CONNECTION_TIMEOUT", strconv.Atoi)
	if err != nil {
		return err
	}

	maxPoolSize, err := envConvertor("DATABASE_MONGO_MAX_POOL_SIZE", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.Database.Mongo.ConnectionOptions.ConnectionTimeout = time.Duration(connectionTimeout * int(time.Minute))
	cfg.Database.Mongo.ConnectionOptions.MaxPoolSize = uint64(maxPoolSize)
	cfg.Database.Mongo.ConnectionOptions.W = os.Getenv("DATABASE_MONGO_MAX_WRITE_CONCERNS")

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

func setSecrets(cfg *AppConfig) {
	cfg.Secrets.EncryptionKey = []byte(os.Getenv("ENCRYPTION_KEY"))
}

func setTokenDetails(cfg *AppConfig) error {
	var atExpiresDays, rtExpiresDays int

	atExpiresDays, err := envConvertor("AT_EXPIRES_DAYS", func(v string) (int, error) {
		return strconv.Atoi(v)
	})
	if err != nil {
		return err
	}

	rtExpiresDays, err = envConvertor("RT_EXPIRES_DAYS", func(v string) (int, error) {
		return strconv.Atoi(v)
	})
	if err != nil {
		return err
	}

	cfg.TokenDetails.AtExpiresDays = atExpiresDays
	cfg.TokenDetails.RtExpiresDays = rtExpiresDays

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
