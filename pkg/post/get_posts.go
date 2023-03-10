package post

import (
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetPosts(c *fiber.Ctx) error {
	return c.JSON("")
}
