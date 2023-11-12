package database

import (
	"convy/conf"
	"database/sql"
	"fmt"
)

func ConnectToPostgres(cfg *conf.AppConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SslMode,
		cfg.Database.Timezone,
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnectionMaxLifetime)

	return sqlDB, err
}
