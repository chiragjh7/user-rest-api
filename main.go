package main

// import the required packages
import (
	"github.com/chiragjh7/user-api/configs"
	"github.com/chiragjh7/user-api/routes"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	// fiber app instance
	app := fiber.New()

	// user-api database connection
	configs.ConnectDB()
	// user-api routes
	routes.UserRoute(app)

	// listen on port from .env file
	var PORT = os.Getenv("PORT")
	app.Listen(PORT)
}
