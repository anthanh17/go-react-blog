package routes

import (
	"github.com/anthanh17/go-react-blog/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, handler *controllers.Handler) {
	app.Post("/api/register", handler.Register)
	app.Post("/api/login", handler.Login)
	app.Get("/api/user", handler.User)
	app.Post("/api/logout", handler.Logout)
}
