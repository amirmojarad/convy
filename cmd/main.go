package main

import (
	"convy/conf"
	"convy/database"
	"convy/internal/logger"
	goose "github.com/pressly/goose/v3"
)

func main() {
	if err := runServer(); err != nil {
		logger.GetLogger().Fatal(err)

		return
	}

}

func runServer() error {
	cfg, err := conf.NewAppConfig()
	if err != nil {
		return err
	}

	sqlDb, err := database.Connect(cfg)
	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err = goose.Up(sqlDb, cfg.Database.MigrationPath); err != nil {
		return err
	}

	return nil
}
