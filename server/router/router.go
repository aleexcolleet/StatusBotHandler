package router

import (
	"MST_FV/config"
	"MST_FV/internal/domain/usecases"
	"MST_FV/server/router/controllers"
	"github.com/gofiber/fiber/v2"
)

/*
SetUpRoutes sets all routes for the server

app.group-> used to group related routes together. This helps to keep
them organized and permits to apply middleware.
*/
func SetUpRoutes(app *fiber.App, msfServices usecases.MSFInterface, cfg config.Config) {

	app.Get("/", rootHandler)
	v1 := app.Group("/api/v1") //base path for api/v1

	//setting public routes
	PublicRouter(v1)

	//setting user routes (consultAndSend, ...)
	userController := controllers.NewUserController(msfServices, cfg)
	UserRouter(v1, userController)
}

// Handler for root endpoint
func rootHandler(c *fiber.Ctx) error {
	return c.SendString(" This is the Fiber Server Root!\n" +
		"You can check public routes:\t/api/v1/public/health\n" +
		"Or the consultAndSend: \n")
}
