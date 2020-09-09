package middleware

import "github.com/gofiber/fiber"

//MaxBody tamanho maximo da request
func MaxBody(size int) fiber.Handler {
	return func(c *fiber.Ctx) {
		if len(c.Body()) >= size {
			c.Next(fiber.ErrRequestEntityTooLarge)
			return
		}
		c.Next()
	}
}
