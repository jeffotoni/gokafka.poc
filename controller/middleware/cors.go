package middleware

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

//Cors middleware
func Cors(app *fiber.App) {
	app.Use(cors.New())
	return
}
