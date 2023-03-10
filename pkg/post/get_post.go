package post

import (
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetPost(c *fiber.Ctx) error {
	return c.JSON("")

}
