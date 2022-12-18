package responses

import "github.com/gofiber/fiber/v2"

// UserResponse is the response structure for user
type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}
