package routes

import (
	"github.com/chiragjh7/user-api/controllers"
	"github.com/gofiber/fiber/v2"
)

// UserRoute is the route for user
func UserRoute(app *fiber.App) {

	app.Get("/user", controllers.GetAllUsers)
	app.Post("/user", controllers.CreateUser)

	// :userId is a parameter that will be passed to the controller to get a specific user
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
}
