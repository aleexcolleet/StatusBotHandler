package router

import (
	"MST_FV/server/router/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(r fiber.Router, controller *controllers.UserController) {

	user := r.Group("/users")
	user.Post("/consult-and-send", controller.CheckAndSend)
}
