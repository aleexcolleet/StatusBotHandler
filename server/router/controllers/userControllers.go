package controllers

import (
	"MST_FV/config"
	"MST_FV/internal/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

/*
UserController is an adapter that handles HTTP requests and delegates
them to UseCases.
*/
type UserController struct {
	UserService usecases.MSFInterface
	cfg         config.Config
}

func NewUserController(userService usecases.MSFInterface, cfg config.Config) *UserController {
	return &UserController{
		UserService: userService,
		cfg:         cfg,
	}
}

func (uc *UserController) CheckAndSend(c *fiber.Ctx) error {

	err := uc.UserService.ConsultAndSend(c.Context(), uc.cfg)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return nil
}
