package middleware

import (
	"github.com/gofiber/fiber"
	mw "github.com/gofiber/fiber/middleware"
)

//Logger log
func Logger(app *fiber.App) {
	app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))
	return
}
