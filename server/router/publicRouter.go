package router

import "github.com/gofiber/fiber/v2"

// PublicRouter define public routes like health checks
func PublicRouter(r fiber.Router) {

	public := r.Group("/public") //Base path for public routes
	health := public.Group("/health")

	health.Get("/", healthChecker)
}

/*
in this case public path would be-> /api/v1/public
*/

func healthChecker(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "Service running correctly",
		"code":     200,
		"success":  true,
		"endpoint": c.Path(),
	})
}
