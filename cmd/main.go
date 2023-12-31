package main

import (
	"convy/conf"
	"convy/database"
	"convy/internal/controller"
	"convy/internal/logger"
	"convy/internal/repository"
	"convy/internal/service"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	goose "github.com/pressly/goose/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	sqlDb, err := database.ConnectToPostgres(cfg)
	if err != nil {
		return err
	}

	mongoClient, err := database.ConnectToMongoDb(cfg)
	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err = goose.Up(sqlDb, cfg.Database.Postgres.MigrationPath); err != nil {
		return err
	}

	routerEngine, err := setupRouter(cfg, sqlDb, mongoClient)
	if err != nil {
		return err
	}

	logger.GetLogger().Infof("running server on port %d ...", cfg.App.Port)

	return routerEngine.Run(fmt.Sprintf(":%d", cfg.App.Port))
}

func setupRouter(cfg *conf.AppConfig, sqlDB *sql.DB, mongoClient *mongo.Client) (*gin.Engine, error) {
	engine := gin.New()

	gormDb, err := getGormDB(cfg, sqlDB)
	if err != nil {
		return nil, err
	}

	v1Group := engine.Group("/v1")
	controller.SetupUserRoutes(v1Group, setupUser(cfg, gormDb))
	controller.SetupUserFollowRoutes(v1Group, setupUserFollow(cfg, gormDb), setupMiddleware(cfg))
	controller.SetupPrivateChatRoutes(v1Group, setupPrivateChat(cfg, gormDb, mongoClient), setupMiddleware(cfg))

	return engine, nil
}

func setupPrivateChat(cfg *conf.AppConfig, gormDb *gorm.DB, mongoClient *mongo.Client) *controller.PrivateChat {
	prChatRepo := repository.NewPrivateChat(gormDb, mongoClient)
	prChatSvc := service.NewPrivateChat(
		cfg,
		logger.GetLogger().WithField("name", "private_chat-service"),
		prChatRepo,
	)
	return controller.NewPrivateChat(
		cfg,
		logger.GetLogger().WithField("name", "private_chat-controller"),
		prChatSvc,
	)
}

func setupUser(cfg *conf.AppConfig, gormDb *gorm.DB) *controller.User {
	userRepository := repository.NewUser(gormDb)
	userSvc := service.NewUser(cfg,
		logger.GetLogger().WithField("name", "user-service"),
		userRepository,
	)

	return controller.NewUser(cfg,
		logger.GetLogger().WithField("name", "user-controller"),
		userSvc,
	)
}

func setupMiddleware(cfg *conf.AppConfig) *controller.Middleware {
	tokenSvc := service.NewToken(cfg)
	return controller.NewMiddleware(tokenSvc)
}

func setupUserFollow(cfg *conf.AppConfig, gormDb *gorm.DB) *controller.UserFollow {
	ufRepository := repository.NewUserFollow(gormDb)

	ufSvc := service.NewUserFollow(cfg,
		logger.GetLogger().WithField("name", "user_follow-service"),
		ufRepository,
	)

	ufCtrl := controller.NewUserFollow(cfg,
		logger.GetLogger().WithField("name", "user_follow-controller"),
		ufSvc,
	)

	return ufCtrl
}

func getGormDB(_ *conf.AppConfig, sqlDB *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
