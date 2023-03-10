package post

import (
	"github.com/gofiber/fiber/v2"
)

func (h handler) DeletePost(c *fiber.Ctx) error {
	return c.JSON("")
}
