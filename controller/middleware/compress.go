package middleware

import (
	"github.com/gofiber/fiber"
	mw "github.com/gofiber/fiber/middleware"
)

//Compress middleware
func Compress(app *fiber.App) {
	app.Use(mw.Compress(mw.CompressLevelBestSpeed))
	return
}
