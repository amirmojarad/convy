package database

import (
	"convy/conf"
	"database/sql"
	"fmt"
)

func ConnectToPostgres(cfg *conf.AppConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Database.Postgres.Host,
		cfg.Database.Postgres.Username,
		cfg.Database.Postgres.Password,
		cfg.Database.Postgres.Name,
		cfg.Database.Postgres.Port,
		cfg.Database.Postgres.SslMode,
		cfg.Database.Postgres.Timezone,
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.Database.Postgres.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.Database.Postgres.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(cfg.Database.Postgres.ConnectionMaxLifetime)

	return sqlDB, err
}
