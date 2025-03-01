package server

import (
	"MST_FV/config"
	"MST_FV/internal/domain/usecases"
	"MST_FV/server/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

type FiberServer struct {
	app         *fiber.App //app instance
	host        string
	port        string
	mSFServices usecases.MSFInterface //repo dependency injection
	config      config.Config         //todo maybe this isn't correct acc to DDD
}

func NewFiberServer(cfg config.Config, msfServices usecases.MSFInterface) *FiberServer {

	app := fiber.New() //creating new fiber web app

	router.SetUpRoutes(app, msfServices, cfg) //todo

	return &FiberServer{
		app:         app,
		host:        cfg.Server.Host,
		port:        cfg.Server.Port,
		mSFServices: msfServices,
		config:      cfg,
	}
}

func (s *FiberServer) Start() {

	log.Printf("🚀 trying to run fiber server at https://%s:%s\n", s.host, s.port)
	err := s.app.Listen(":" + s.port)
	if err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
}
