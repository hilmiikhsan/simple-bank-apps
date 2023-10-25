package delivery

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/simple-bank-apps/config"
	"github.com/simple-bank-apps/config/database/redis"
	"github.com/simple-bank-apps/delivery/controller/auth"
	"github.com/simple-bank-apps/delivery/controller/payment"
	"github.com/simple-bank-apps/manager"
	"github.com/simple-bank-apps/middleware"
	"github.com/sirupsen/logrus"
)

type appServer struct {
	engine         *gin.Engine
	config         *config.Config
	useCaseManager manager.UsecaseManager
	middleware     middleware.TokenMiddleware
	logMiddleware  middleware.LogMiddleware
}

func (a *appServer) initController() {
	auth.NewAuthController(a.engine, a.useCaseManager.AuthUsecase(), a.middleware, a.logMiddleware)
	payment.NewPaymentController(a.engine, a.useCaseManager.PaymentUsecase(), a.middleware, a.logMiddleware)
}

func (a *appServer) Run() {
	a.initController()
	err := a.engine.Run(a.config.App.Port)
	if err != nil {
		log.Println("Error running server :", err.Error())
	}
}

func Server() *appServer {
	engine := gin.Default()
	log := logrus.New()

	err := config.LoadConfig("./env.yaml")
	if err != nil {
		log.Println("Error loading config file :", err.Error())
	}

	infraManager, err := manager.NewInfraManager(config.Cfg)
	if err != nil {
		log.Println("Error init infra manager :", err.Error())
	}

	redisServer := redis.NewRedisConfig(config.Cfg)
	rdb, err := redisServer.Connect(context.Background())
	if err != nil {
		log.Fatalln("error config redis", err.Error())
	}

	logger := middleware.NewLogMiddleware(log, config.Cfg)
	jwt := middleware.NewJWT(config.Cfg, rdb)
	tokenMiddleware := middleware.NewTokenMiddleware(jwt)
	repositoryManager := manager.NewRepositoryManager(infraManager)
	useCaseManager := manager.NewUsecaseManager(repositoryManager, jwt, config.Cfg, logger)

	return &appServer{
		engine:         engine,
		config:         config.Cfg,
		useCaseManager: useCaseManager,
		middleware:     tokenMiddleware,
		logMiddleware:  logger,
	}
}
