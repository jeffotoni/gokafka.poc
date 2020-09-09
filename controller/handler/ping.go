package handler

import "github.com/gofiber/fiber"

//Ping pong
func Ping(c *fiber.Ctx) {
	c.Status(200).Send("Pong 🏓")
}
