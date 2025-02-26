package server

import (
	"MST_FV/config"
	"MST_FV/internal/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app            *fiber.App //app instance
	cfg            config.Config
	consultAndSend usecases.ConsultAndSendUseCase //repo dependency injection
}

func NewServer(cfg config.Config, consultAndSend usecases.ConsultAndSendUseCase) *Server {
	return &Server{
		app:            fiber.New(),
		cfg:            cfg,
		consultAndSend: consultAndSend,
	}
}

func (s *Server) SetUpRoutes() {

	//Root endpoint
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	//ConsultAndSendUseCase call
	s.app.Post("/consultAndSend", s.HandleConsultAndSend)
	//Healt endpoint
	s.app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Status OK")
	})
}
func (s *Server) HandleConsultAndSend(ctx *fiber.Ctx) error {
	err := s.consultAndSend.ConsultAndSend(ctx.Context(), s.cfg)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{ // internal server error mapped
			"error": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "consultAndSend executed successfully",
		"url":     "/consultAndSend",
	})
}

func (s *Server) Start() error {
	return s.app.Listen(":3000")
}
